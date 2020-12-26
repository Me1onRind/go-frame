package interceptor

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"go-frame/global"
	"google.golang.org/grpc/metadata"
)

func getRequestID(ctx context.Context) string {
	md, _ := metadata.FromIncomingContext(ctx)
	values := md.Get(global.ProtocolRequestIDKey)

	if len(values) == 0 {
		return uuid.NewV4().String()
	}
	return values[0]
}

func getJWTToken(ctx context.Context) string {
	md, _ := metadata.FromIncomingContext(ctx)
	values := md.Get(global.ProtocolJWTTokenKey)
	if len(values) == 0 {
		return ""
	}
	return values[0]
}
