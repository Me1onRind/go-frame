package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"go-frame/global"
	"go-frame/internal/pkg/context"
	"go-frame/internal/pkg/errcode"
	"go-frame/internal/pkg/logger"
	"go-frame/internal/utils/auth"
	"go-frame/internal/utils/encode"
	"time"
)

type JWTService struct {
}

func NewJWTService() *JWTService {
	return &JWTService{}
}

func (j *JWTService) GenerateToken(ctx context.Context, appKey, appSecret string) (string, int64, *errcode.Error) {
	nowTime := time.Now()
	jwtSetting := global.JWTSetting
	expireAt := nowTime.Add(jwtSetting.Expire).Unix()

	claims := &auth.Claims{
		AppKey:    encode.MD5(appKey),
		AppSecret: encode.MD5(appSecret),
		StandardClaims: jwt.StandardClaims{
			Issuer:    jwtSetting.Issuser,
			ExpiresAt: expireAt,
		},
	}

	logger.WithTrace(ctx).WithFields(
		logger.JSONKV("claims", claims), logger.KV("secret", jwtSetting.Secret),
	).Info("jwt signed")
	tokenClainms := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClainms.SignedString([]byte(jwtSetting.Secret))
	if err != nil {
		return "", 0, errcode.JWTSignedFailError.WithError(err)
	}
	return token, expireAt, nil
}
