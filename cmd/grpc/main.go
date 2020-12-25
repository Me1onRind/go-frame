package main

import (
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"go-frame/global"
	"go-frame/internal/controller/grpc/user"
	"go-frame/internal/initialize"
	"go-frame/internal/pkg/interceptor"
	"go-frame/internal/pkg/setting"
	"go-frame/proto/pb"
	"google.golang.org/grpc"
	"net"
)

func init() {
	global.Environment = initialize.InitEnvironment()
	if err := SetGrpcSetting(); err != nil {
		panic(err)
	}
	if err := initialize.SetupStore(); err != nil {
		panic(err)
	}
	if err := initialize.SetupLogger(); err != nil {
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
	}

	for k, v := range LoadSections {
		if err := st.ReadSection(k, v); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			interceptor.Logger,
			interceptor.Recover,
		)),
	}

	s := grpc.NewServer(opts...)
	pb.RegisterUserServiceServer(s, user.NewUserGrpcController())

	st := global.GrpcServerSetting
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", st.Host, st.Port))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Start grpc server %s\n", fmt.Sprintf("%s:%d", st.Host, st.Port))
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
