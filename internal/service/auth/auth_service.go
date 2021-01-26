package auth

import (
	"go-frame/internal/constant/auth_constant"
	"go-frame/internal/core/custom_ctx"
	"go-frame/internal/core/errcode"
	authDao "go-frame/internal/dao/auth"
	"go-frame/internal/lib/auth"
	"time"
)

type AuthService struct {
	AuthDao *authDao.AuthDao
}

func NewAuthService() *AuthService {
	return &AuthService{
		AuthDao: authDao.NewAuthDao(),
	}
}

func (a *AuthService) GenerateJWTToken(ctx *custom_ctx.Context, appKey, appSecret string) (string, int64, *errcode.Error) {
	authInfo, err := a.AuthDao.GetAuthByAppKey(ctx, appKey)
	if err != nil {
		return "", 0, err
	}

	if authInfo == nil || authInfo.Config == nil {
		return "", 0, errcode.RecordNotFound.WithInfo("Auth or Auth config not exist")
	}

	if authInfo.AppSecret != appSecret {
		return "", 0, errcode.AppSecretWrongError
	}

	param := &auth.GenerateJwtTokenParam{
		AppKey:    appKey,
		AppSecret: appSecret,
	}

	if authInfo.Config.Flag&auth_constant.ExpireSwitch != 0 {
		param.Expires = time.Duration(authInfo.Config.Expires) * time.Second
	}

	return auth.GenerateJwtToken(ctx, param)
}

func (a *AuthService) ListAuths(ctx *custom_ctx.Context) ([]*authDao.Auth, int64, *errcode.Error) {
	return a.AuthDao.ListAuths(ctx, 0, 10000)
}
