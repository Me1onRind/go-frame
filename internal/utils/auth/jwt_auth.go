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

func GenerateToken(tracer logger.Tracer, appKey, appSecret string) (string, int64, *errcode.Error) {
	nowTime := time.Now()
	jwtSetting := global.JWTSetting
	expireAt := nowTime.Add(jwtSetting.Expire).Unix()

	claims := &Claims{
		AppKey:    encode.MD5(appKey),
		AppSecret: encode.MD5(appSecret),
		StandardClaims: jwt.StandardClaims{
			Issuer:    jwtSetting.Issuser,
			ExpiresAt: expireAt,
		},
	}

	logger.WithTrace(tracer).WithFields(
		logger.JSONKV("claims", claims), logger.KV("secret", jwtSetting.Secret),
	).Info("jwt signed")
	tokenClainms := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClainms.SignedString([]byte(jwtSetting.Secret))
	if err != nil {
		return "", 0, errcode.JWTSignedFailError.WithError(err)
	}
	return token, expireAt, nil
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
