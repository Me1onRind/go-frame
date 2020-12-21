package global

import (
	"gorm.io/gorm"
)

const (
	DefaultDB = "default"
)

var (
	ReadDBs  = map[string]*gorm.DB{}
	WriteDBs = map[string]*gorm.DB{}
)
