package auth

import (
	"go-frame/internal/pkg/context"
	"go-frame/internal/pkg/errcode"
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

func (a *AuthController) GenerateToken(ctx *context.HttpContext) (interface{}, *errcode.Error) {
	request := ctx.Raw.(*auth_proto.GenerateTokenReq)
	token, expireAt, err := a.AuthService.GenerateJWTToken(ctx, request.AppKey, request.AppSecret)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"token":     token,
		"expire_at": expireAt,
	}, nil
}

func (a *AuthController) ListAuths(ctx *context.HttpContext) (interface{}, *errcode.Error) {
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
