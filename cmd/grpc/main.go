package main

import (
	"fmt"
	//micro "github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/server"
	svr "github.com/micro/go-micro/v2/server/grpc"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"go-frame/global"
	"go-frame/internal/controller/grpc/user"
	"go-frame/internal/initialize"
	"go-frame/internal/pkg/setting"
	"go-frame/internal/pkg/wrapper"
	"go-frame/proto/pb"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	global.Environment = initialize.InitEnvironment()
	if err := SetGrpcSetting(); err != nil {
		panic(err)
	}
	if err := initialize.SetupZapLogger(); err != nil {
		panic(err)
	}
	if err := initialize.SetupStore(); err != nil {
		panic(err)
	}
	if err := initialize.SetupJaegerTracer("go-frame-grpc"); err != nil {
		panic(err)
	}
}

func SetGrpcSetting() error {
	st, err := setting.NewSetting(global.Environment.Env, "./configs/", "yml")
	if err != nil {
		return err
	}

	LoadSections := map[string]interface{}{
		"GrpcServer":   &global.GrpcServerSetting,
		"Mysql":        &global.MysqlSetting,
		"Logger.Info":  &global.InfoLoggerSetting,
		"Logger.Error": &global.ErrorLoggerSetting,
		"JWT":          &global.JWTSetting,
		"Etcd":         &global.EtcdSetting,
	}

	for k, v := range LoadSections {
		if err := st.ReadSection(k, v); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	st := global.GrpcServerSetting
	address := fmt.Sprintf("%s:%d", st.Host, st.Port)

	grpcServer := svr.NewServer(
		server.Name(global.GrpcServerSetting.Name),
		server.Version("latest"),
		server.Address(address),
		server.Registry(
			etcdv3.NewRegistry(
				registry.Addrs(global.EtcdSetting.Addresses...),
				registry.Timeout(st.RegistryTimeout),
			),
		),
		server.WrapHandler(wrapper.Tracing),
		server.WrapHandler(wrapper.AccessLogger),
		server.WrapHandler(wrapper.JWT),
		server.WrapHandler(wrapper.ErrHandler),
	)

	pb.RegisterUserServiceHandler(grpcServer, user.NewUserGrpcController())
	if err := grpcServer.Start(); err != nil {
		panic(err)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-sigs:
	}

	grpcServer.Stop()

}
