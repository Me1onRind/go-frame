package user

type User struct {
	ID       uint64 `gorm:"column:id;primaryKey"`
	UserID   uint64 `gorm:"column:user_id"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	CTime    uint32 `gorm:"column:ctime"`
	MTime    uint32 `gorm:"column:mtime"`
}

func (u *User) TableName() string {
	return "user_tab"
}
