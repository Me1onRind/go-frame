package global

import (
	"go-frame/internal/pkg/setting"
)

var (
	Environment        *setting.EnvironmentS
	HttpServerSetting  *setting.HttpServerSettingS
	GrpcServerSetting  *setting.GrpcServerSettingS
	MysqlSetting       *setting.DBSettingS
	InfoLoggerSetting  *setting.LoggerSettingS
	ErrorLoggerSetting *setting.LoggerSettingS
)
