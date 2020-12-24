package initialize

import (
	"fmt"
	"github.com/Me1onRind/logrotate"
	"github.com/gorilla/sessions"
	"go-frame/global"
	"go-frame/internal/pkg/logger"
	"go-frame/internal/pkg/setting"
	"go-frame/internal/pkg/store"
	"io"
	"path/filepath"
	"strings"
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

func SetupLogger() error {
	infoWriter, err := initLogWriter(global.InfoLoggerSetting)
	if err != nil {
		return err
	}
	errorWriter, err := initLogWriter(global.ErrorLoggerSetting)
	if err != nil {
		return err
	}

	loggerConfig := &logger.Config{
		NewWriter: func(level logger.Level) io.Writer {
			if level <= logger.LevelInfo {
				return infoWriter
			}
			return errorWriter
		},
		TimeFormat: logger.DefaultTimeFormat,
		Level:      logger.LevelInfo,
		Separate:   '|',
	}
	logger.SetLogger(logger.NewLogger(loggerConfig))
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
