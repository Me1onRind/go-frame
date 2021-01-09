package wrapper

import (
	"context"
	"github.com/micro/go-micro/v2/server"
	"go-frame/global"
	customContext "go-frame/internal/core/context"
	"go-frame/internal/utils/ctx_helper"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
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

		if spanCtx == nil {
			t := span.SpanContext()
			spanCtx = &t
		}
		spanID = spanCtx.SpanID.String()
		traceID = spanCtx.TraceID.String()

		commonCtx := customContext.NewCommonContext(c,
			customContext.WithSpan(span),
			customContext.WithZapLogger(
				global.Logger.With(
					zap.String("traceID", traceID),
					zap.String("spanID", spanID),
				),
			),
		)
		c = ctx_helper.SetCommonContext(c, commonCtx)

		err := fn(c, req, resp)

		return err
	}
}
