package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-frame/global"
	"go-frame/internal/initialize"
	"go-frame/internal/initialize/validation"
	"go-frame/internal/pkg/setting"
	"go-frame/internal/routers"
)

func init() {
	global.Environment = initialize.InitEnvironment()
	if err := SetupHttpSetting(); err != nil {
		panic(err)
	}
	if err := initialize.SetupStore(); err != nil {
		panic(err)
	}
	if err := initialize.SetupLogger(); err != nil {
		panic(err)
	}
	if err := validation.RegisterGinValidation(); err != nil {
		panic(err)
	}
	if err := initialize.SetupCookie(); err != nil {
		panic(err)
	}
	if err := initialize.SetupJaegerTracer("go-frame-api"); err != nil {
		panic(err)
	}
	initialize.InitGrpcClient()
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
	r.Run(fmt.Sprintf("%s:%d", httpServerSetting.Host, httpServerSetting.Port))
}
