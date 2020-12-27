package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	opentracing "github.com/opentracing/opentracing-go"
	"go-frame/global"
)

func Tracing() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx context.Context
		span := opentracing.SpanFromContext(c.Request.Context())
		if span != nil {
			span, ctx = opentracing.StartSpanFromContextWithTracer(c.Request.Context(), global.Tracer,
				c.Request.URL.Path, opentracing.ChildOf(span.Context()))
		} else {
			span, ctx = opentracing.StartSpanFromContextWithTracer(c.Request.Context(), global.Tracer,
				c.Request.URL.Path)
		}
		defer span.Finish()

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
