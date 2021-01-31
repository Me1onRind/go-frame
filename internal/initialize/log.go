package initialize

import (
	"go-frame/global"
	"go-frame/internal/core/logger"
	"go-frame/internal/core/setting"
	"path/filepath"
	"strings"
	"time"

	"github.com/Me1onRind/logrotate"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

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
