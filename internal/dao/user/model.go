package user

import (
	"go-frame/internal/lib/base"
)

type User struct {
	*base.Model
	UserID   uint64 `gorm:"column:user_id"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
}

func (u *User) TableName() string {
	return "user"
}
