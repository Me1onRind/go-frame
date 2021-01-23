package context

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-frame/global"
	"go-frame/internal/core/errcode"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type TranscationFunc func() *errcode.Error

type Context struct {
	GinCtx *gin.Context
	Ctx    context.Context

	Span   trace.Span
	Logger *zap.Logger
	Env    string

	txs     map[string]*gorm.DB
	traceID string
}

func NewContext(logger *zap.Logger) *Context {
	return &Context{
		Logger: logger,
		txs:    map[string]*gorm.DB{},
	}
}

func (c *Context) ReadDB(dbKey string) *gorm.DB {
	if db := c.txs[dbKey]; db != nil {
		return db
	}

	if db := global.ReadDBs[dbKey]; db != nil {
		return db
	}

	panic(fmt.Sprintf("Can't get read db, dbKey[%s]", dbKey))
}

func (c *Context) WriteDB(dbKey string) *gorm.DB {
	if db := c.txs[dbKey]; db != nil {
		return db
	}

	if db := global.WriteDBs[dbKey]; db != nil {
		return db
	}

	panic(fmt.Sprintf("Can't get write db, dbKey[%s]", dbKey))
}

func (c *Context) Transaction(dbKey string, fc TranscationFunc) (err *errcode.Error) {
	// allow nested
	if db := c.txs[dbKey]; db != nil {
		return fc()
	}

	tx := c.WriteDB(dbKey).Begin()
	c.txs[dbKey] = tx

	paniced := true
	defer func() {
		if paniced || err != nil {
			tx.Rollback()
		}
		c.txs[dbKey] = nil
	}()

	err = fc()
	if err == nil {
		if e := tx.Commit().Error; e != nil {
			err = errcode.DBError.WithError(e)
		}
	}

	paniced = false
	return err
}

func (c *Context) TraceID() string {
	return c.traceID
}

func (c *Context) SetLoggerPrefix(fields ...zap.Field) {
	c.Logger = c.Logger.With(fields...)
}
