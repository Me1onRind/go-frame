package interceptor

import (
	"context"
	newContext "go-frame/internal/pkg/context"
	"go-frame/internal/utils/auth"
	"google.golang.org/grpc"
)

func JWT(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	token := getJWTToken(ctx)

	newCtx := ctx.(*newContext.CommonContext)
	if err := auth.JWTAuth(newCtx, token); err != nil {
		return nil, err.ToRpcError()
	}

	resp, err := handler(ctx, req)
	return resp, err
}
