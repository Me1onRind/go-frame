package interceptor

import (
	"context"
	newContext "go-frame/internal/pkg/context"
	"go-frame/internal/pkg/errcode"
	"go-frame/internal/pkg/logger"
	"google.golang.org/grpc"
	"runtime/debug"
)

func Recover(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		newCtx := ctx.(*newContext.CommonContext)
		if e := recover(); e != nil {
			logger.WithTrace(newCtx).Errorf("Panic recover err:%v, stack:\n%s", e, debug.Stack())
			resp = nil
			err = errcode.ServerError.ToRpcError()
		}
	}()
	resp, err = handler(ctx, req)
	return resp, err
}
