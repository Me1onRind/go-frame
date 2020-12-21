package model

type Group struct {
	*Model
	GroupName string `gorm:"column:group_name"`
}
