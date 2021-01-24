package initialize

import (
	"fmt"
	"go-frame/global"
	"go-frame/internal/core/logger"
	"go-frame/internal/core/setting"
	"go-frame/internal/core/store"
	"os"
	"path/filepath"
	"strings"
	"time"

	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmot"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	"github.com/Me1onRind/logrotate"
	"github.com/gorilla/sessions"
	opentracing "github.com/opentracing/opentracing-go"
	"go.opentelemetry.io/otel/exporters/trace/jaeger"
	"go.opentelemetry.io/otel/label"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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

func SetupLogger(setGoMicroLogger bool) error {
	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder,
		TimeKey:     "ts",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		CallerKey:        "file",
		EncodeCaller:     zapcore.ShortCallerEncoder,
		ConsoleSeparator: "|",
		EncodeDuration:   zapcore.MillisDurationEncoder,
	})

	initLogWriter := func(loggerSetting *setting.LoggerSettingS) (*logrotate.RotateLog, error) {
		sep := string([]byte{filepath.Separator})
		loggerSetting.LogDir = strings.Trim(loggerSetting.LogDir, sep)
		logPath := loggerSetting.LogDir + sep + loggerSetting.LogName
		writer, err := logrotate.NewRoteteLog(logPath+".2006010215",
			logrotate.WithRotateTime(global.InfoLoggerSetting.RotateTimeDuration),
			logrotate.WithCurLogLinkname(logPath),
			logrotate.WithDeleteExpiredFile(global.InfoLoggerSetting.MaxAge, loggerSetting.LogName+".*"),
		)
		if err != nil {
			return nil, err
		}
		return writer, nil
	}

	infoWriter, err := initLogWriter(global.InfoLoggerSetting)
	if err != nil {
		return err
	}
	warnWriter, err := initLogWriter(global.ErrorLoggerSetting)
	if err != nil {
		return err
	}

	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.WarnLevel
	})
	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel
	})
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(warnWriter), warnLevel),
	)

	global.Logger = zap.New(core, zap.AddCaller())
	if setGoMicroLogger {
		logger.SetGoMicroLogger(global.Logger.WithOptions(zap.AddCallerSkip(2)))
	}

	return nil
}

func SetupJaegerTracer(serviceName string) error {
	var err error
	global.JaegerPipelineFlush, err = jaeger.InstallNewPipeline(
		jaeger.WithCollectorEndpoint("http://localhost:14268/api/traces"),
		jaeger.WithProcess(jaeger.Process{
			ServiceName: serviceName,
			Tags: []label.KeyValue{
				label.String("exporter", "jaeger"),
			},
		}),
		jaeger.WithSDK(&sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()}),
	)
	return err
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
