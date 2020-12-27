package user

import (
	"context"
	newContext "go-frame/internal/pkg/context"
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

func (u *UserGrpcController) GetUserInfo(c context.Context, req *pb.GetUserReq, resp *pb.UserInfo) error {
	ctx := newContext.GetCommonContext(c)
	userInfo, err := u.UserService.GetUserByUserID(ctx, req.UserID, true)
	if err != nil {
		return err.ToRpcError()
	}
	resp.UserID = userInfo.UserID
	resp.Username = userInfo.Username
	return nil
}
