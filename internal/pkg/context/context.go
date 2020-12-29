package context

import (
	"context"
	"fmt"
	"go-frame/global"
	"go-frame/internal/pkg/errcode"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type TranscationFunc func() *errcode.Error

type Context interface {
	context.Context

	Env() string

	ReadDB(dbKey string) *gorm.DB
	WriteDB(dbKey string) *gorm.DB
	Transaction(dbKey string, fc TranscationFunc) *errcode.Error

	Span() trace.Span
	Logger() *zap.Logger
}

type contextS struct {
	txs    map[string]*gorm.DB
	span   trace.Span
	logger *zap.Logger
}

func newContextS() *contextS {
	return &contextS{
		txs: map[string]*gorm.DB{},
	}
}

func (c *contextS) Env() string {
	return global.Environment.Env
}

func (c *contextS) ReadDB(dbKey string) *gorm.DB {
	if db := c.txs[dbKey]; db != nil {
		return db
	}

	if db := global.ReadDBs[dbKey]; db != nil {
		return db
	}

	panic(fmt.Sprintf("Can't get read db, dbKey[%s]", dbKey))
}

func (c *contextS) WriteDB(dbKey string) *gorm.DB {
	if db := c.txs[dbKey]; db != nil {
		return db
	}

	if db := global.WriteDBs[dbKey]; db != nil {
		return db
	}

	panic(fmt.Sprintf("Can't get write db, dbKey[%s]", dbKey))
}

func (c *contextS) Transaction(dbKey string, fc TranscationFunc) (err *errcode.Error) {
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

func (c *contextS) Span() trace.Span {
	return c.span
}

func (c *contextS) Logger() *zap.Logger {
	return c.logger
}
