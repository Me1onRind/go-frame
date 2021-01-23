package wrapper

import (
	"context"
	customCtx "go-frame/internal/core/context"
	"go-frame/internal/utils/optracing"

	"github.com/micro/go-micro/v2/server"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

func Tracing(fn server.HandlerFunc) server.HandlerFunc {
	return func(c context.Context, req server.Request, resp interface{}) error {
		carrier := optracing.NewOpentracingCarrierFromGrpcContext(c)
		//carrier := propagation.TextMapCarrier(md)
		c = otel.GetTextMapPropagator().Extract(c, carrier)

		tr := otel.Tracer("tracer")
		_, span := tr.Start(c, req.Method())
		defer span.End()

		ctx := customCtx.GetFromContext(c)
		spanCtx := span.SpanContext()

		ctx.Span = span
		ctx.SetLoggerPrefix(zap.String("TraceID", spanCtx.TraceID.String()), zap.String("spanID", spanCtx.SpanID.String()))

		return fn(c, req, resp)
	}
}
