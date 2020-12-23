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
}

func SetupHttpSetting() error {
	st, err := setting.NewSetting(global.Environment.Env, "./configs/", "yml")
	if err != nil {
		return err
	}
	if err := st.ReadSection("HttpServer", &global.HttpServerSetting); err != nil {
		return err
	}
	if err := st.ReadSection("Mysql", &global.MysqlSetting); err != nil {
		return err
	}
	if err := st.ReadSection("Logger.Info", &global.InfoLoggerSetting); err != nil {
		return err
	}
	if err := st.ReadSection("Logger.Error", &global.ErrorLoggerSetting); err != nil {
		return err
	}

	return nil
}

func main() {
	r := routers.NewRouter()
	httpServerSetting := global.HttpServerSetting
	gin.SetMode(httpServerSetting.RunMode)
	r.Run(fmt.Sprintf("%s:%d", httpServerSetting.Host, httpServerSetting.Port))
}
