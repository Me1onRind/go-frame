package global

import (
	etcd "github.com/coreos/etcd/clientv3"
	"github.com/micro/go-micro/v2/client"
	minio "github.com/minio/minio-go"
)

var (
	GrpcClient  client.Client
	MinioClient *minio.Client
	EtcdClient  *etcd.Client
)
