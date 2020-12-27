package wrapper

import (
	"context"
	"github.com/micro/go-micro/v2/server"
	"go-frame/internal/lib/auth"
	newContext "go-frame/internal/pkg/context"
)

func JWT(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		token := getJWTToken(ctx)
		newCtx := newContext.GetCommonContext(ctx)
		if err := auth.JWTAuth(newCtx, token); err != nil {
			return err.ToRpcError()
		}

		return fn(ctx, req, resp)
	}
}
