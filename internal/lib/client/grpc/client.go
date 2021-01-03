package grpc

import (
	"context"
	"github.com/micro/go-micro/v2/client"
	cli "github.com/micro/go-micro/v2/client/grpc"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"go-frame/global"
	customContext "go-frame/internal/core/context"
	"go-frame/internal/utils/ctx_helper"
	"go.uber.org/zap"
	"time"
)

var (
	GoFrameClient client.Client
)

func InitClients() {
	GoFrameClient = newGoFrameClient()
}

func JWTContext(ctx customContext.Context, jwtToken string) context.Context {
	newCtx := ctx_helper.SetCustomContext(context.Background(), ctx)
	return metadata.NewContext(newCtx, map[string]string{
		global.ProtocolJWTTokenKey: jwtToken,
		global.ProtocolSpanIDKey:   ctx.Span().SpanContext().SpanID.String(),
		global.ProtocolTraceIDKey:  ctx.Span().SpanContext().TraceID.String(),
	})
}

func newGoFrameClient() client.Client {
	return cli.NewClient(
		client.Registry(
			etcdv3.NewRegistry(
				registry.Addrs(global.EtcdSetting.Addresses...),
				registry.Timeout(5*time.Second),
			),
		),
		client.WrapCall(callLogger),
	)
}

func callLogger(fn client.CallFunc) client.CallFunc {
	return func(ctx context.Context, node *registry.Node, req client.Request, rsp interface{}, opts client.CallOptions) error {
		c := ctx_helper.GetCustomContext(ctx)
		begin := time.Now()
		err := fn(ctx, node, req, rsp, opts)

		c.Logger().Info("GRPC Call End",
			zap.String("method", req.Method()),
			zap.Any("reqBody", req.Body()),
			zap.String("target", node.Address),
			zap.Any("resp", rsp),
			zap.Error(err),
			zap.Duration("cost", time.Since(begin)),
		)

		return err
	}
}
