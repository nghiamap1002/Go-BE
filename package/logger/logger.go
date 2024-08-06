package logger

import (
	"os"
	"personal/ShopDev/Go-BE/package/setting"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(config setting.LoggerSetting) *LoggerZap {

	logLevel := config.LogLevel
	// debug => info => warn => error => fatal => panic

	var level zapcore.Level

	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	// case "fatal":
	// 	level = zapcore.FatalLevel
	// case "panic":
	// 	level = zapcore.PanicLevel
	default:
		level = zapcore.InfoLevel
	}

	encoder := getEncoderLog()

	hook := lumberjack.Logger{
		Filename:   config.FileLogName,
		MaxSize:    config.MaxSize, // megabytes
		MaxBackups: config.MaxBackup,
		MaxAge:     config.MaxAge,   //days
		Compress:   config.Compress, // disabled by default
	}

	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		level)

	return &LoggerZap{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))}
}

func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	// 1722247062.19373 => 2024-07-29T16:57:42.193+0700
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// ts => time
	encodeConfig.TimeKey = "time"

	//  from info INFO
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encodeConfig)
}
