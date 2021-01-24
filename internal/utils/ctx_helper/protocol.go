package ctx_helper

import (
	"context"
	"go-frame/global"
	customContext "go-frame/internal/core/context"

	"github.com/micro/go-micro/v2/metadata"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

func JWTContext(ctx *customContext.Context, jwtToken string) context.Context {
	carrier := opentracing.TextMapCarrier{}
	if err := opentracing.GlobalTracer().Inject(ctx.Span().Context(), opentracing.TextMap, &carrier); err != nil {
		ctx.Logger().Warn("Extract span fail", zap.Error(err))
	}

	carrier[global.ProtocolJWTTokenKey] = jwtToken
	carrier[global.ProtocolRequestID] = ctx.RequestID()

	return metadata.NewContext(ctx, map[string]string(carrier))
}
