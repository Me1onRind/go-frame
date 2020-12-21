package store

import (
	"fmt"
	"go-frame/internal/pkg/setting"
	"go-frame/internal/utils/date"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

	db.Callback().Create().Before("*").Register("my_plugin:create_ctime_mtime", createCTimeAndMTime)

	db.Callback().Update().Before("*").Register("my_plugin:update_mtime", updateMTime)

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
