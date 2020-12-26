package auth

import (
	"github.com/dgrijalva/jwt-go"
	"go-frame/global"
	"go-frame/internal/pkg/errcode"
	"go-frame/internal/pkg/logger"
	"go-frame/internal/utils/encode"
	"time"
)

type Claims struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	jwt.StandardClaims
}

type GenerateJwtTokenParam struct {
	AppKey    string
	AppSecret string
	Expires   time.Duration
}

func GenerateJwtToken(tracer logger.Tracer, param *GenerateJwtTokenParam) (string, int64, *errcode.Error) {
	jwtSestting := global.JWTSetting
	claims := &Claims{
		AppKey:    encode.MD5(param.AppKey),
		AppSecret: encode.MD5(param.AppSecret),
		StandardClaims: jwt.StandardClaims{
			Issuer: jwtSestting.Issuer,
		},
	}

	if param.Expires > 0 {
		nowTime := time.Now()
		claims.StandardClaims.ExpiresAt = nowTime.Add(time.Duration(param.Expires) * time.Second).Unix()
	}

	logger.WithTrace(tracer).WithFields(
		logger.JSONKV("claims", claims), logger.KV("secret", jwtSestting.Secret),
	).Info("jwt signed")
	tokenClainms := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClainms.SignedString([]byte(jwtSestting.Secret))
	if err != nil {
		return "", 0, errcode.JWTSignedFailError.WithError(err)
	}
	return token, claims.ExpiresAt, nil
}

func JWTAuth(tracer logger.Tracer, token string) *errcode.Error {
	logger.WithTrace(tracer).Infof("jwt token:%s", token)
	_, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.JWTSetting.Secret), nil
	})

	if err != nil {
		switch err.(*jwt.ValidationError).Errors {
		case jwt.ValidationErrorExpired:
			return errcode.JWTTimeoutError.WithError(err)
		default:
			return errcode.JWTAuthorizedFailError.WithError(err)
		}
	}

	return nil
}