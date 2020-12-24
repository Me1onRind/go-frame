package global

import (
	"github.com/gorilla/sessions"
)

var (
	CookieStore *sessions.CookieStore
)

const (
	UserInfoSessionKey = "accout"
	UserSessionName    = "user"
)
