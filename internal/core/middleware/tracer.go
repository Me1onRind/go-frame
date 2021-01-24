package middleware

import (
	"go-frame/global"
	"go-frame/internal/core/context"
	"go-frame/internal/core/errcode"
	"go-frame/internal/core/gateway"
	"go-frame/internal/utils/optracing"

	"github.com/gin-gonic/gin"
	opentracing "github.com/opentracing/opentracing-go"
	"go.elastic.co/apm/module/apmhttp"
	"go.uber.org/zap"
)

func Tracing() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.Request.Header.Get(global.ProtocolRequestID)
		traceParent := c.Request.Header.Get(apmhttp.W3CTraceparentHeader)

		if len(traceParent) == 0 && len(requestID) > 0 {
			c.Request.Header.Set(apmhttp.W3CTraceparentHeader, optracing.RequestIDToTraceparent(requestID))
		} else if len(traceParent) > 0 && len(requestID) == 0 {
			requestID = optracing.RequestIDFromW3CTraceparent(traceParent)
		}

		spanCtx, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, c.Request.Header)
		if err != nil && err != opentracing.ErrSpanContextNotFound {
			c.JSON(200, gateway.NewResponse(errcode.OptExtractError, nil))
			c.Abort()
			return
		}

		span := opentracing.StartSpan(c.Request.URL.Path, opentracing.ChildOf(spanCtx))
		defer span.Finish()

		if len(requestID) == 0 {
			requestID = optracing.RequestIDFromSpan(span.Context())
		}

		ctx := context.GetFromGinContext(c)
		ctx.SetSpan(span)
		ctx.SetRequestID(requestID)
		ctx.SetLoggerPrefix(zap.String("requestID", requestID))

		c.Next()
	}
}
