package global

import (
	"go-frame/internal/core/setting"
)

var (
	Environment        *setting.EnvironmentS
	HttpServerSetting  *setting.HttpServerSettingS
	GrpcServerSetting  *setting.GrpcServerSettingS
	MysqlSetting       *setting.DBSettingS
	InfoLoggerSetting  *setting.LoggerSettingS
	ErrorLoggerSetting *setting.LoggerSettingS
	JWTSetting         *setting.JWTSettingS
	EtcdSetting        *setting.EtcdSettingS
	MinioSetting       *setting.MinioSettingS
	RedisSetting       *setting.RedisSettings
)
