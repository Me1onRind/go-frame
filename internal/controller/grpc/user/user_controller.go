package user

import (
	"context"
	"go-frame/internal/service/user"
	"go-frame/proto/pb"
)

type UserGrpcController struct {
	UserService *user.UserService
}

func NewUserGrpcController() *UserGrpcController {
	return &UserGrpcController{
		UserService: user.NewUserService(),
	}
}

func (u *UserGrpcController) GetUserInfo(ctx context.Context, request *pb.GetUserRequest) (*pb.UserInfo, error) {
	return &pb.UserInfo{
		UserID:   1,
		Username: "username",
	}, nil
}
