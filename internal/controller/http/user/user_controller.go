package user

import (
	"go-frame/internal/core/custom_ctx"
	"go-frame/internal/core/errcode"
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

func (u *UserController) Login(ctx *custom_ctx.Context, raw interface{}) (interface{}, *errcode.Error) {
	request := raw.(*user_proto.LoginReq)
	token, err := u.UserService.Login(ctx, request.Username, request.Password)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"token": token,
	}, nil
}

func (u *UserController) GetUserInfo(ctx *custom_ctx.Context, raw interface{}) (interface{}, *errcode.Error) {
	request := raw.(*user_proto.GetUserInfoReq)
	userInfo, err := u.UserService.GetUserFromRemote(ctx, request.UserID)
	if err != nil {
		return nil, err
	}

	return userInfo, nil
}

func (u *UserController) GetUserInfoByToken(ctx *custom_ctx.Context, raw interface{}) (interface{}, *errcode.Error) {
	request := raw.(*user_proto.GetUserInfoByTokenReq)
	userInfo, err := u.UserService.GetUserInfoByToken(ctx, request.Token)
	if err != nil {
		return nil, err
	}

	return userInfo, nil
}
