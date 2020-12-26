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

func (u *UserGrpcController) GetUserInfo(ctx context.Context, req *pb.GetUserReq, resp *pb.UserInfo) error {
	resp.UserID = 1
	resp.Username = "username"
	return nil
}
