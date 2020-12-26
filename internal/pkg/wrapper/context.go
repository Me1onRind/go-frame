package wrapper

import (
	"context"
	"github.com/micro/go-micro/v2/server"
	newContext "go-frame/internal/pkg/context"
)

func NewContext(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		requestID := getRequestID(ctx)
		newCtx := newContext.NewCommonContext(ctx, newContext.WithRequestID(requestID))
		return fn(newCtx, req, resp)
	}
}
