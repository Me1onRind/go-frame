package main

import (
	"fmt"
	"go-frame/global"
	"go-frame/internal/core/setting"
	"go-frame/internal/initialize"
	"go-frame/internal/routers"

	"github.com/gin-gonic/gin"
)

func init() {
	global.Environment = initialize.InitEnvironment()
	if err := SetupHttpSetting(); err != nil {
		panic(err)
	}
	if err := initialize.SetupLogger(false); err != nil {
		panic(err)
	}
	if err := initialize.SetupStore(); err != nil {
		panic(err)
	}
	if err := initialize.SetupCookie(); err != nil {
		panic(err)
	}
	if err := initialize.RegisterGinValidation(); err != nil {
		panic(err)
	}
	if err := initialize.SetClients(); err != nil {
		panic(err)
	}
	initialize.SetupOpentracingTracer()
}

func SetupHttpSetting() error {
	st, err := setting.NewSetting(global.Environment.Env, "./configs/", "yml")
	if err != nil {
		return err
	}

	LoadSections := map[string]interface{}{
		"HttpServer":   &global.HttpServerSetting,
		"Mysql":        &global.MysqlSetting,
		"Logger.Info":  &global.InfoLoggerSetting,
		"Logger.Error": &global.ErrorLoggerSetting,
		"JWT":          &global.JWTSetting,
		"Etcd":         &global.EtcdSetting,
		"Minio":        &global.MinioSetting,
		"Redis":        &global.RedisSetting,
	}

	for k, v := range LoadSections {
		if err := st.ReadSection(k, v); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	r := routers.NewRouter()
	httpServerSetting := global.HttpServerSetting
	gin.SetMode(httpServerSetting.RunMode)
	_ = r.Run(fmt.Sprintf("%s:%d", httpServerSetting.Host, httpServerSetting.Port))
}
