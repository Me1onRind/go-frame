package user

import (
	"go-frame/internal/core/context"
	"go-frame/internal/core/errcode"
	"go-frame/internal/core/session"
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

func (u *UserController) Login(ctx *context.Context, raw interface{}) (interface{}, *errcode.Error) {
	request := raw.(*user_proto.LoginReq)
	userInfo, err := u.UserService.Login(ctx, request.Username, request.Password)
	if err != nil {
		return nil, err
	}

	sessionUserInfo := &session.UserInfo{
		UserID:   userInfo.ID,
		Username: userInfo.Username,
	}
	if err := session.SetUserInfo(ctx, sessionUserInfo); err != nil {
		return nil, err
	}

	return sessionUserInfo, nil
}

func (u *UserController) GetUserInfo(ctx *context.Context, raw interface{}) (interface{}, *errcode.Error) {
	request := raw.(*user_proto.GetUserInfoReq)
	userInfo, err := u.UserService.GetUserFromRemote(ctx, request.UserID)
	if err != nil {
		return nil, err
	}

	return userInfo, nil
}
