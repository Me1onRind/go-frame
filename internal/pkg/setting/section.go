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
	Host             string
	Port             uint32
	Name             string
	RegistryTimeout  time.Duration
	RegistryInterVal time.Duration
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

type JWTSettingS struct {
	Issuer string
	Secret string
}

type EtcdSettingS struct {
	Addresses []string
}
