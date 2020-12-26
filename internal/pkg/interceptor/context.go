package interceptor

import (
	"context"
	newContext "go-frame/internal/pkg/context"
	"google.golang.org/grpc"
)

func NewContext(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	requestID := getRequestID(ctx)
	newCtx := newContext.NewCommonContext(ctx, newContext.WithRequestID(requestID))

	resp, err := handler(newCtx, req)
	return resp, err
}
