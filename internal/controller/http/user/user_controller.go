package user

import (
	"go-frame/internal/lib/session"
	"go-frame/internal/pkg/context"
	"go-frame/internal/pkg/errcode"
	"go-frame/internal/service/user"
	"go-frame/protocol/user_proto"
)

type UserController struct {
	UserService *user.UserService
}

func NewUserContoller() *UserController {
	return &UserController{
		UserService: user.NewUserService(),
	}
}

func (u *UserController) Login(ctx *context.HttpContext) (interface{}, *errcode.Error) {
	request := ctx.Raw.(*user_proto.LoginReq)
	userInfo, err := u.UserService.Login(ctx, request.Username, request.Password)
	if err != nil {
		return nil, err
	}

	sessionUserInfo := &session.UserInfo{
		UserID:   userInfo.ID,
		Username: userInfo.Username,
	}
	session.SetUserInfo(ctx.Context, sessionUserInfo)

	return sessionUserInfo, nil
}

func (u *UserController) GetUserInfo(ctx *context.HttpContext) (interface{}, *errcode.Error) {
	request := ctx.Raw.(*user_proto.GetUserInfoReq)
	userInfo, err := u.UserService.GetUserFromRemote(ctx, request.UserID)
	if err != nil {
		return nil, err
	}

	return userInfo, nil
}
