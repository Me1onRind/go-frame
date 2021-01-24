package initialize

import (
	"context"
	"go-frame/global"
	customCtx "go-frame/internal/core/context"
	"time"

	"github.com/micro/go-micro/v2/client"
	cli "github.com/micro/go-micro/v2/client/grpc"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	minio "github.com/minio/minio-go"
	"go.uber.org/zap"
)

func SetClients() error {
	global.GrpcClient = newGrpcClient(global.EtcdSetting.Addresses)

	// minio client
	//var err error
	//minioSetting := global.MinioSetting
	//global.MinioClient, err = newMinioClient(minioSetting.Endpoint, minioSetting.AccessKeyID, minioSetting.SecretAccessKey)
	//if err != nil {
	////return err
	//}
	return nil
}

func newGrpcClient(addresses []string) client.Client {
	return cli.NewClient(
		client.Registry(
			etcd.NewRegistry(
				registry.Addrs(addresses...),
				registry.Timeout(5*time.Second),
			),
		),
		client.WrapCall(callLogger),
	)
}

func callLogger(fn client.CallFunc) client.CallFunc {
	return func(ctx context.Context, node *registry.Node, req client.Request, rsp interface{}, opts client.CallOptions) error {
		c := customCtx.GetFromContext(ctx)
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

func newMinioClient(endpoint, accessKeyID, secreAccessKey string) (*minio.Client, error) {
	return minio.New(endpoint, accessKeyID, secreAccessKey, false)
}
