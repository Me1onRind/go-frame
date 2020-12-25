package interceptor

import (
	"context"
	"go-frame/internal/pkg/errcode"
	"go-frame/internal/pkg/logger"
	"google.golang.org/grpc"
	"runtime/debug"
)

func Recover(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		if e := recover(); e != nil {
			tracer := getTracer(ctx)
			logger.WithTrace(tracer).Errorf("Panic recover err:%v, stack:\n%s", e, debug.Stack())
			resp = nil
			err = errcode.ServerError.ToRpcError()
		}
	}()
	resp, err = handler(ctx, req)
	return resp, err
}
