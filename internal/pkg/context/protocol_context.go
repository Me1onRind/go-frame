package context

import (
	"context"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"go-frame/global"
)

type HttpContext struct {
	Raw interface{}

	*contextS
	*gin.Context

	reqeustID string
}

func NewHttpContext(c *gin.Context) *HttpContext {
	h := &HttpContext{
		contextS: newContextS(),
		Context:  c,
	}
	h.contextS.reqeustID = c.GetString(global.ContextRequestIDKey)
	return h
}

func (h *HttpContext) RequestID() string {
	return h.reqeustID
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
