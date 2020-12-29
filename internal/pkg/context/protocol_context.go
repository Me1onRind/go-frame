package context

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-frame/internal/pkg/logger"
	"go.opentelemetry.io/otel/trace"
)

type HttpContext struct {
	Raw interface{}

	*contextS
	*gin.Context
}

func NewHttpContext(c *gin.Context, opts ...Option) *HttpContext {
	h := &HttpContext{
		contextS: newContextS(),
		Context:  c,
	}

	for _, opt := range opts {
		opt(h.contextS)
	}

	return h
}

type CommonContext struct {
	*contextS
	context.Context
}

func NewCommonContext(ctx context.Context, opts ...Option) *CommonContext {
	c := &CommonContext{
		contextS: newContextS(),
		Context:  ctx,
	}

	for _, opt := range opts {
		opt(c.contextS)
	}

	return c
}

type Option func(ctx *contextS)

func WithSpan(span trace.Span) Option {
	return func(ctx *contextS) {
		ctx.span = span
	}
}

func WithTracer(tracer logger.Tracer) Option {
	return func(ctx *contextS) {
		ctx.Tracer = tracer
	}
}
