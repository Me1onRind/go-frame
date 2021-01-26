package ctx_helper

import (
	"context"
	"go-frame/internal/constant/proto_constant"
	customCtx "go-frame/internal/core/custom_ctx"

	"github.com/micro/go-micro/v2/metadata"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

func JWTContext(ctx *customCtx.Context, jwtToken string) context.Context {
	carrier := opentracing.TextMapCarrier{}
	if err := opentracing.GlobalTracer().Inject(ctx.Span().Context(), opentracing.TextMap, &carrier); err != nil {
		ctx.Logger().Warn("Extract span fail", zap.Error(err))
	}

	carrier[proto_constant.ProtocolJWTTokenKey] = jwtToken
	carrier[proto_constant.ProtocolRequestID] = ctx.RequestID()

	return metadata.NewContext(ctx, map[string]string(carrier))
}
