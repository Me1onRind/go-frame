package store

import (
	"fmt"
	"go-frame/internal/core/custom_ctx"
	"go-frame/internal/core/setting"
	"go-frame/internal/utils/date"

	"github.com/opentracing/opentracing-go"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	spanKey = "span"
)

func NewDBEngine(dbSetting *setting.DBSettingS) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local&timeout=%dms&readTimeout=%dms&writeTimeout=%dms",
		dbSetting.Username, dbSetting.Password,
		dbSetting.Host, dbSetting.Port,
		dbSetting.DBName, dbSetting.ConnectTimeout.Milliseconds(),
		dbSetting.ReadTimeout.Milliseconds(), dbSetting.WriteTimeout.Milliseconds())

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(dbSetting.MaxIdleConns)
	sqlDB.SetMaxOpenConns(dbSetting.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(dbSetting.ConnectMaxLifeTime)

	_ = db.Callback().Query().Before("*").Register("my_plugin:tracing_start", tracingStart())
	_ = db.Callback().Query().After("*").Register("my_plugin:tracing_end", tracingEnd("SELECT"))
	_ = db.Callback().Create().Before("*").Register("my_plugin:create_ctime_mtime", createCTimeAndMTime)
	_ = db.Callback().Update().Before("*").Register("my_plugin:update_mtime", updateMTime)

	return db, nil
}

func updateMTime(db *gorm.DB) {
	db.Statement.SetColumn("mtime", date.UnixTime())
}

func createCTimeAndMTime(db *gorm.DB) {
	now := date.UnixTime()
	db.Statement.SetColumn("ctime", now)
	db.Statement.SetColumn("mtime", now)
}

func tracingStart() func(*gorm.DB) {
	return func(db *gorm.DB) {
		if db.Statement.Context != nil {
			if ctx, ok := db.Statement.Context.(*custom_ctx.Context); ok {
				span := ctx.Span().Tracer().StartSpan("db", opentracing.ChildOf(ctx.Span().Context()))
				db.Set(spanKey, span)
			}
		}
	}
}

func tracingEnd(operation string) func(*gorm.DB) {
	return func(db *gorm.DB) {
		if sp, ok := db.Get(spanKey); ok {
			span := sp.(opentracing.Span)
			defer span.Finish()
			span.SetTag("db.table", db.Statement.Table)
			span.SetTag("db.method", operation)
			span.SetTag("db.rowAffected", db.RowsAffected)
			span.SetTag("db.err", db.Error)
		}
	}
}
