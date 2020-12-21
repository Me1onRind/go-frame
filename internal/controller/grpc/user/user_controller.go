package user

import (
	"context"
	"go-frame/proto/pb"
)

type UserGrpcController struct{}

func NewUserGrpcController() *UserGrpcController {
	return &UserGrpcController{}
}

func (u *UserGrpcController) GetUserInfo(ctx context.Context, request *pb.GetUserRequest) (*pb.UserInfo, error) {
	return &pb.UserInfo{
		UserID:   1,
		Username: "username",
	}, nil
}
