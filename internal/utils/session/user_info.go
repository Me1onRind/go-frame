package session

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	//"github.com/gorilla/sessions"
	"go-frame/global"
)

const (
	UserInfoKey = "account"
)

type UserInfo struct {
	UserID   uint64 `json:"userID"`
	Username string `json:"username"`
}

func SetUserInfo(c *gin.Context, userInfo *UserInfo) {
	session, _ := global.CookieStore.Get(c.Request, global.UserSessionName)
	userInfoBytes, _ := json.Marshal(userInfo)
	session.Values[UserInfoKey] = userInfoBytes
	session.Save(c.Request, c.Writer)
}

func GetUserInfo(c *gin.Context) *UserInfo {
	session, _ := global.CookieStore.Get(c.Request, global.UserSessionName)
	if value, ok := session.Values[UserInfoKey]; ok {
		if userInfoBytes, ok := value.([]byte); ok {
			var userInfo UserInfo
			_ = json.Unmarshal(userInfoBytes, &userInfo)
			return &userInfo
		}
	}
	return nil
}
