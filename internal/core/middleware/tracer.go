package middleware

import (
	"github.com/gin-gonic/gin"
	"go-frame/global"
	customContext "go-frame/internal/core/context"
	"go-frame/internal/utils/ctx_helper"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
	//"go.opentelemetry.io/otel/trace"
	"context"
)

func Tracing() gin.HandlerFunc {
	return func(c *gin.Context) {

		tr := otel.Tracer("tracer")
		_, span := tr.Start(context.Background(), c.Request.URL.Path)
		defer span.End()

		traceID := span.SpanContext().TraceID.String()
		spanID := span.SpanContext().SpanID.String()

		httpContext := customContext.NewHttpContext(c,
			customContext.WithSpan(span),
			customContext.WithZapLogger(
				global.Logger.With(
					zap.String("traceID", traceID),
					zap.String("spanID", spanID),
				),
			),
		)
		ctx_helper.SetHttpContext(c, httpContext)

		c.Next()
	}
}
