package middleware

import (
	"github.com/gin-gonic/gin"
	customContext "go-frame/internal/pkg/context"
	"go-frame/internal/pkg/logger"
	"go-frame/internal/utils/ctx_helper"
	"go.opentelemetry.io/otel"
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

		var tracer logger.Tracer = logger.NewSimpleTracer(
			logger.KV("traceID", traceID),
			logger.KV("spanID", spanID),
		)

		httpContext := customContext.NewHttpContext(c, customContext.WithTracer(tracer), customContext.WithSpan(span))
		ctx_helper.SetHttpContext(c, httpContext)

		c.Next()
	}
}
