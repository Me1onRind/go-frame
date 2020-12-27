package context

import (
	"context"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"go-frame/global"
)

const (
	saveInContextKey = "commonCtx"
)

type HttpContext struct {
	Raw interface{}

	*contextS
	*gin.Context
}

func NewHttpContext(c *gin.Context) *HttpContext {
	h := &HttpContext{
		contextS: newContextS(),
		Context:  c,
	}
	h.contextS.reqeustID = c.GetString(global.ContextRequestIDKey)
	return h
}

type CommonContext struct {
	*contextS
	context.Context
}

type CommonCtxOption func(ctx *CommonContext)

func NewCommonContext(ctx context.Context, opts ...CommonCtxOption) *CommonContext {
	c := &CommonContext{
		contextS: newContextS(),
		Context:  ctx,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func SetCommonContext(c context.Context, ctx *CommonContext) context.Context {
	return context.WithValue(c, saveInContextKey, ctx)
}

func GetCommonContext(c context.Context) *CommonContext {
	return c.Value(saveInContextKey).(*CommonContext)
}

func WithRequestID(reqeustID string) CommonCtxOption {
	return func(ctx *CommonContext) {
		ctx.contextS.reqeustID = reqeustID
	}
}

func WithAutoRequestID() CommonCtxOption {
	return func(ctx *CommonContext) {
		ctx.contextS.reqeustID = uuid.NewV4().String()
	}
}
