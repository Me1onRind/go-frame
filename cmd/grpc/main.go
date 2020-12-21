package main

import (
	"fmt"
	"go-frame/global"
	"go-frame/internal/controller/grpc/user"
	"go-frame/proto/pb"
	"google.golang.org/grpc"
	"net"
)

func main() {
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, user.NewUserGrpcController())

	st := global.GrpcServerSetting
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", st.Host, st.Port))
	if err != nil {
		panic(err)
	}

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
