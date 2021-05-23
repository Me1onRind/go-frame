package global

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

const (
	DefaultDB = "default"
)

var (
	ReadDBs  = map[string]*gorm.DB{}
	WriteDBs = map[string]*gorm.DB{}

	Redis *redis.Client
)
