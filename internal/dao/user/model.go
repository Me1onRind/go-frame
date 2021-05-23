package user

import (
	"go-frame/internal/utils/slice"
)

type User struct {
	ID       uint64            `gorm:"column:id;primaryKey" json:"id"`
	UserID   uint64            `gorm:"column:user_id" json:"userId"`
	Username string            `gorm:"column:username" json:"username"`
	Password string            `gorm:"column:password" json:"password"`
	RoleIDs  slice.Uint32Slice `gorm:"column:role_ids" json:"role_ids"`
	CTime    uint32            `gorm:"column:ctime" json:"ctime"`
	MTime    uint32            `gorm:"column:mtime" json:"mtime"`
}

type CacheUserInfo struct {
	UserID    uint64   `json:"user_id"`
	Username  string   `json:"username"`
	RoleNames []string `json:"roles"`
}

func (u *User) TableName() string {
	return "user_tab"
}
