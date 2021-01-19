package global

import (
	"github.com/micro/go-micro/v2/client"
	minio "github.com/minio/minio-go"
)

var (
	GrpcClient  client.Client
	MinioClient *minio.Client
)
