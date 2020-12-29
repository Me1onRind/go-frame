package initialize

import (
	"fmt"
	"github.com/Me1onRind/logrotate"
	"github.com/gorilla/sessions"
	"go-frame/global"
	"go-frame/internal/pkg/setting"
	"go-frame/internal/pkg/store"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"path/filepath"
	"strings"
	"time"
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

func SetupZapLogger() error {
	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder,
		TimeKey:     "ts",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		CallerKey:    "file",
		EncodeCaller: zapcore.ShortCallerEncoder,
	})

	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.WarnLevel
	})
	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel
	})

	infoWriter, err := initLogWriter(global.InfoLoggerSetting)
	if err != nil {
		return err
	}
	warnWriter, err := initLogWriter(global.ErrorLoggerSetting)
	if err != nil {
		return err
	}

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(warnWriter), warnLevel),
	)

	global.Logger = zap.New(core, zap.AddCaller())

	return nil
}

func initLogWriter(loggerSetting *setting.LoggerSettingS) (*logrotate.RotateLog, error) {
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
