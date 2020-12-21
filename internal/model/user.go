package model

type User struct {
	*Model
	Username string `gorm:"column:username"`
	Passwd   string `gorm:"column:passwd"`
	GroupId  uint64 `gorm:"column:group_id"`
}

func (u *User) TableName() string {
	return "user"
}
