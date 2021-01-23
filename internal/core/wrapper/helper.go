package wrapper

import (
	"context"
	"encoding/hex"
	"github.com/micro/go-micro/v2/metadata"
	"go-frame/global"
	"go.opentelemetry.io/otel/trace"
)

func getJWTToken(ctx context.Context) string {
	jwtToken, _ := metadata.Get(ctx, global.ProtocolJWTTokenKey)
	return jwtToken
}

func getSpanCtx(ctx context.Context) *trace.SpanContext {
	return nil
}

func stringToTraceID(str string) trace.TraceID {
	data, _ := hex.DecodeString(str)
	var b [16]byte = [16]byte{}
	for k, v := range data[:16] {
		b[k] = v
	}
	return b
}

func stringToSpanID(str string) trace.SpanID {
	data, _ := hex.DecodeString(str)
	var b [8]byte = [8]byte{}
	for k, v := range data[:8] {
		b[k] = v
	}
	return b
}
