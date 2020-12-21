package main

import (
	"fmt"
	"go-frame/global"
	"go-frame/internal/controller/grpc/user"
	"go-frame/internal/initialize"
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
}

func SetGrpcSetting() error {
	st, err := setting.NewSetting(global.Environment.Env, "./configs/", "yml")
	if err != nil {
		return err
	}
	if err := st.ReadSection("GrpcServer", &global.GrpcServerSetting); err != nil {
		return err
	}
	if err := st.ReadSection("Mysql", &global.MysqlSetting); err != nil {
		return err
	}
	if err := st.ReadSection("Logger.Info", &global.InfoLoggerSetting); err != nil {
		return err
	}
	if err := st.ReadSection("Logger.Error", &global.ErrorLoggerSetting); err != nil {
		return err
	}

	return nil
}

func main() {
	s := grpc.NewServer()
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
