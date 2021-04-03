package zap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"testing"
	"time"
)

// 本文参考:https://www.jianshu.com/p/7349b6227233

func NewEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "1",
		LevelKey:       "1",
		NameKey:        "1",
		CallerKey:      "1",
		MessageKey:     "1",
		StacktraceKey:  "1",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func TestLogrotate(t *testing.T) {
	// 配置写入文件，需要：lumberjack支持
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "zap_demo2_test.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
	})
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(NewEncoderConfig()),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), w),
		zap.DebugLevel,
	)
	logger := zap.New(core, zap.AddCaller())

	//2019-09-26 14:34:47.991	INFO	zap/zap_demo2_test.go:47	info
	logger.Info("info")
	logger.Debug("hello debug")
}
