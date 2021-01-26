package wrapper

import (
	"context"
	"go-frame/internal/constant/proto_constant"
	customCtx "go-frame/internal/core/custom_ctx"
	"go-frame/internal/utils/optracing"

	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/server"
	opentracing "github.com/opentracing/opentracing-go"
	"go.elastic.co/apm/module/apmhttp"
	"go.uber.org/zap"
)

func Tracing(fn server.HandlerFunc) server.HandlerFunc {
	return func(c context.Context, req server.Request, resp interface{}) error {
		md, _ := metadata.FromContext(c)
		requestID := md[proto_constant.ProtocolRequestID]
		traceParent := md[apmhttp.W3CTraceparentHeader]
		if len(traceParent) == 0 && len(requestID) > 0 {
			md[apmhttp.W3CTraceparentHeader] = optracing.RequestIDToTraceparent(requestID)
		} else if len(traceParent) > 0 && len(requestID) == 0 {
			requestID = optracing.RequestIDFromW3CTraceparent(traceParent)
		}

		carrier := opentracing.TextMapCarrier(md)
		spanCtx, _ := opentracing.GlobalTracer().Extract(opentracing.TextMap, carrier)

		var span opentracing.Span
		if spanCtx != nil {
			span = opentracing.StartSpan(req.Method(), opentracing.ChildOf(spanCtx))
		} else {
			span = opentracing.StartSpan(req.Method())
		}
		defer span.Finish()

		if len(requestID) == 0 {
			requestID = optracing.RequestIDFromSpan(span.Context())
		}

		ctx := customCtx.GetFromContext(c)
		ctx.SetSpan(span)
		ctx.SetRequestID(requestID)
		ctx.SetLoggerPrefix(zap.String("requestID", requestID))

		return fn(c, req, resp)
	}
}
