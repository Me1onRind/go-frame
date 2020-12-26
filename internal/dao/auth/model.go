package auth

import (
	"go-frame/internal/lib/base"
)

type Auth struct {
	*base.Model
	AppName   string      `gorm:"column:app_name"`
	AppKey    string      `gorm:"column:app_key"`
	AppSecret string      `gorm:"column:app_secret"`
	ConfigID  uint64      `gorm:"config_id"`
	Config    *AuthConfig `gorm:"foreignkey:id;references:config_id"`
}

func (a *Auth) TableName() string {
	return "auth"
}

type AuthConfig struct {
	*base.Model
	ConfigName string `gorm:"column:config_name"`
	Expires    uint32 `gorm:"column:expires"`
	Flag       uint64 `gorm:"column:flag"`
}

func (a *AuthConfig) TableName() string {
	return "auth_config"
}
