package wrapper

import (
	"context"
	"go-frame/global"
	customCtx "go-frame/internal/core/context"
	"go-frame/internal/core/errcode"
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
		requestID := md[global.ProtocolRequestID]
		traceParent := md[apmhttp.W3CTraceparentHeader]
		if len(traceParent) == 0 && len(requestID) > 0 {
			md[apmhttp.W3CTraceparentHeader] = optracing.RequestIDToTraceparent(requestID)
		} else if len(traceParent) > 0 && len(requestID) == 0 {
			requestID = optracing.RequestIDFromW3CTraceparent(traceParent)
		}

		carrier := opentracing.TextMapCarrier(md)
		spanCtx, err := opentracing.GlobalTracer().Extract(opentracing.TextMap, carrier)
		if err != nil && err != opentracing.ErrSpanContextNotFound {
			return errcode.OptExtractError.ToRpcError()
		}
		span := opentracing.StartSpan(req.Method(), opentracing.ChildOf(spanCtx))
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
