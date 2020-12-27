package initialize

import (
	"go-frame/internal/lib/client/grpc"
)

func InitGrpcClient() {
	grpc.InitClient()
}
