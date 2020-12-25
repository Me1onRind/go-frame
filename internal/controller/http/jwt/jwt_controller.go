package jwt

import (
	"go-frame/internal/pkg/context"
	"go-frame/internal/pkg/errcode"
	"go-frame/internal/service/jwt"
	"go-frame/protocol/jwt_proto"
)

type JWTController struct {
	JWTService *jwt.JWTService
}

func NewJWTController() *JWTController {
	return &JWTController{}
}

func (j *JWTController) GenerateToken(ctx *context.HttpContext) (interface{}, *errcode.Error) {
	request := ctx.Raw.(*jwt_proto.GenerateTokenReq)
	token, expireAt, err := j.JWTService.GenerateToken(ctx, request.AppKey, request.AppSecret)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"token":     token,
		"expire_at": expireAt,
	}, nil
}
