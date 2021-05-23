package initialize

import (
	"fmt"
	"go-frame/global"
	"go-frame/internal/core/store"
	"os"

	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmot"

	"github.com/gorilla/sessions"
	opentracing "github.com/opentracing/opentracing-go"
)

func SetupStore() error {
	writeDB, err := store.NewDBEngine(global.MysqlSetting)
	if err != nil {
		return err
	}
	global.WriteDBs[global.DefaultDB] = writeDB

	readDB, err := store.NewDBEngine(global.MysqlSetting)
	if err != nil {
		return err
	}
	global.ReadDBs[global.DefaultDB] = readDB

	global.Redis = store.NewRedisClientPool(global.RedisSetting)

	return nil
}

func SetupCookie() error {
	cookiesSetting := global.HttpServerSetting.Cookies
	if cookiesSetting.StoreType == "CookieStore" {
		global.CookieStore = sessions.NewCookieStore([]byte(cookiesSetting.SecretKey))
	} else {
		return fmt.Errorf("Unsupport storeType:%s", cookiesSetting.StoreType)
	}
	return nil
}

func SetupOpentracingTracer() {
	os.Setenv("ELASTIC_APM_SERVER_URL", "http://localhost:8200")
	os.Setenv("ELASTIC_APM_SECRET_TOKEN", "")
	os.Setenv("ELASTIC_APM_STACK_TRACE_LIMIT", "0")
	os.Setenv("ELASTIC_APM_USE_ELASTIC_TRACEPARENT_HEADER", "false")
	os.Setenv("ELASTIC_APM_TRANSACTION_SAMPLE_RATE", "1.0")
	//os.Setenv("ELASTIC_APM_DISABLE_METRICS", "system.process.*")
	tracer, _ := apm.NewTracer("go-frame", "0.0.1")
	opentracing.SetGlobalTracer(apmot.New(apmot.WithTracer(tracer)))
}
