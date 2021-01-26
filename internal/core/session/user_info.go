package session

import (
	"encoding/json"
	"go-frame/internal/constant/proto_constant"
	"go-frame/internal/core/custom_ctx"

	"github.com/gin-gonic/gin"

	//"github.com/gorilla/sessions"
	"go-frame/global"
	"go-frame/internal/core/errcode"
)

const (
	UserInfoKey = "account"
)

type UserInfo struct {
	UserID   uint64 `json:"userID"`
	Username string `json:"username"`
}

func SetUserInfo(c *custom_ctx.Context, userInfo *UserInfo) *errcode.Error {
	session, _ := global.CookieStore.Get(c.GinCtx.Request, proto_constant.UserSessionName)
	userInfoBytes, _ := json.Marshal(userInfo)
	session.Values[UserInfoKey] = userInfoBytes
	if err := session.Save(c.GinCtx.Request, c.GinCtx.Writer); err != nil {
		return errcode.SaveSessionError.WithError(err)
	}
	return nil
}

func GetUserInfo(c *gin.Context) *UserInfo {
	session, _ := global.CookieStore.Get(c.Request, proto_constant.UserSessionName)
	if value, ok := session.Values[UserInfoKey]; ok {
		if userInfoBytes, ok := value.([]byte); ok {
			var userInfo UserInfo
			_ = json.Unmarshal(userInfoBytes, &userInfo)
			return &userInfo
		}
	}
	return nil
}
