package grpc

import (
	"context"
	micro "github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"go-frame/global"
	customContext "go-frame/internal/pkg/context"
	"go-frame/proto/pb"
	"time"
)

var (
	GoFrameClient pb.UserService
)

func InitClient() {
	GoFrameClient = pb.NewUserService("go-frame-grpc", initServiceClient("go-frame-grpc.client"))
}

func JwtContext(ctx customContext.Context, jwtToken string) context.Context {
	return metadata.NewContext(context.Background(), map[string]string{
		global.ProtocolRequestIDKey: ctx.RequestID(),
		global.ProtocolJWTTokenKey:  jwtToken,
	})
}

func initServiceClient(microName string) client.Client {
	service := micro.NewService(
		micro.Name(microName),
		micro.RegisterInterval(5*time.Second),
		micro.Registry(etcdv3.NewRegistry(
			registry.Addrs(global.EtcdSetting.Addresses...),
			registry.Timeout(5*time.Second),
		)),
	)
	service.Init()
	/*cli := client.NewClient()*/
	return service.Client()
}
