package user

import (
	"go-frame/internal/utils/base"
)

type User struct {
	*base.Model
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
}

func (u *User) TableName() string {
	return "user"
}
