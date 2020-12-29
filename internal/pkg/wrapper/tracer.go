package wrapper

import (
	"context"
	"github.com/micro/go-micro/v2/server"
	customContext "go-frame/internal/pkg/context"
	"go-frame/internal/pkg/logger"
	"go-frame/internal/utils/ctx_helper"
	//"go.opentelemetry.io/oteltest"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func Tracing(fn server.HandlerFunc) server.HandlerFunc {
	return func(c context.Context, req server.Request, resp interface{}) error {
		var spanID, traceID string
		spanCtx := getSpanCtx(c)
		var ctx context.Context

		if spanCtx != nil {
			ctx = trace.ContextWithRemoteSpanContext(context.Background(), *spanCtx)
		} else {
			ctx = context.Background()
		}

		tr := otel.Tracer("tracer")
		_, span := tr.Start(ctx, req.Method())
		defer span.End()

		spanID = spanCtx.SpanID.String()
		traceID = spanCtx.TraceID.String()

		var tracer logger.Tracer = logger.NewSimpleTracer(
			logger.KV("traceID", traceID),
			logger.KV("spanID", spanID),
		)

		commonCtx := customContext.NewCommonContext(c, customContext.WithTracer(tracer), customContext.WithSpan(span))
		c = ctx_helper.SetCommonContext(c, commonCtx)

		return fn(c, req, resp)
	}
}
