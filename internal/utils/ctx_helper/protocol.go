package ctx_helper

import (
	"context"
	"go-frame/global"
	customContext "go-frame/internal/core/context"

	"github.com/micro/go-micro/v2/metadata"
)

func JWTContext(ctx customContext.Context, jwtToken string) context.Context {
	return metadata.NewContext(context.Background(), map[string]string{
		global.ProtocolJWTTokenKey: jwtToken,
		global.ProtocolSpanIDKey:   ctx.Span().SpanContext().SpanID.String(),
		global.ProtocolTraceIDKey:  ctx.Span().SpanContext().TraceID.String(),
	})
}
