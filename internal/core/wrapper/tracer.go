package wrapper

import (
	"context"
	"fmt"
	customCtx "go-frame/internal/core/context"
	"go-frame/internal/utils/optracing"

	"github.com/micro/go-micro/v2/server"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

func Tracing(fn server.HandlerFunc) server.HandlerFunc {
	return func(c context.Context, req server.Request, resp interface{}) error {
		//carrier := optracing.NewOpentracingCarrierFromGrpcContext(c)
		//c = otel.GetTextMapPropagator().Extract(c, carrier)
		carrier := optracing.NewOpentracingCarrier()

		tr := otel.Tracer("tracer")
		_, span := tr.Start(c, req.Method())
		defer span.End()

		ctx := customCtx.GetFromContext(c)
		spanCtx := span.SpanContext()
		otel.GetTextMapPropagator().Inject(trace.ContextWithSpan(ctx, span), carrier)
		fmt.Println(carrier)

		ctx.Span = span
		ctx.SetLoggerPrefix(zap.String("traceID", spanCtx.TraceID.String()))

		return fn(c, req, resp)
	}
}
