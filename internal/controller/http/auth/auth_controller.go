package auth

import (
	"go-frame/internal/core/context"
	"go-frame/internal/core/errcode"
	"go-frame/internal/service/auth"
	"go-frame/protocol"
	"go-frame/protocol/auth_proto"
)

type AuthController struct {
	AuthService *auth.AuthService
}

func NewAuthController() *AuthController {
	return &AuthController{
		AuthService: auth.NewAuthService(),
	}
}

func (a *AuthController) GenerateToken(ctx *context.Context, raw interface{}) (interface{}, *errcode.Error) {
	request := raw.(*auth_proto.GenerateTokenReq)
	token, expireAt, err := a.AuthService.GenerateJWTToken(ctx, request.AppKey, request.AppSecret)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"token":     token,
		"expire_at": expireAt,
	}, nil
}

func (a *AuthController) ListAuths(ctx *context.Context, raw interface{}) (interface{}, *errcode.Error) {
	list, total, err := a.AuthService.ListAuths(ctx)
	if err != nil {
		return nil, err
	}

	return &protocol.ListResp{
		Page:     0,
		PageSize: 10000,
		Total:    total,
		List:     list,
	}, nil
}
