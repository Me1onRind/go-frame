package context

import (
	"context"
	"fmt"
	"go-frame/global"
	"go-frame/internal/core/errcode"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type TranscationFunc func() *errcode.Error

type Context struct {
	context.Context
	GinCtx *gin.Context

	Span   trace.Span
	Logger *zap.Logger
	Env    string

	requestID string
	txs       map[string]*gorm.DB
	span      opentracing.Span
}

func NewContext(logger *zap.Logger, libCtx context.Context) *Context {
	c := &Context{
		Logger:  logger,
		txs:     map[string]*gorm.DB{},
		Context: libCtx,
	}
	if libCtx == nil {
		c.Context = LoadIntoContext(c, context.Background())
	}
	return c
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

func (c *Context) SetRequestID(requestID string) {
	c.requestID = requestID
}

func (c *Context) SetLoggerPrefix(fields ...zap.Field) {
	c.Logger = c.Logger.With(fields...)
}

func (c *Context) SetSpan(span opentracing.Span) {
	c.span = span
}

func (c *Context) RequestID() string {
	return c.requestID
}
