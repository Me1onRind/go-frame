package setting

import (
	"time"
)

type HttpServerSettingS struct {
	RunMode string
	Host    string
	Port    uint32
	Cookies struct {
		StoreType string
		SecretKey string
	}
}

type GrpcServerSettingS struct {
	Host string
	Port uint32
}

type DBSettingS struct {
	Host               string
	Port               uint32
	Username           string
	Password           string
	DBName             string
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	ConnectTimeout     time.Duration
	MaxIdleConns       int
	MaxOpenConns       int
	ConnectMaxLifeTime time.Duration
}

type LoggerSettingS struct {
	LogDir             string
	LogName            string
	RotateTimeDuration time.Duration
	MaxAge             time.Duration
}
