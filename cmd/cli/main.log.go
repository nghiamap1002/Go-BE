package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// sugar := zap.NewExample().Sugar()
	// sugar.Infof("hello name: %s, age: %d", "concobebe", 20)
	// logger := zap.NewExample()
	// logger.Info("abc", zap.String("name", "abco"), zap.Int("age", 20))

	// logger := zap.NewExample()
	// logger.Info("NewExample")

	// logger, _ = zap.NewDevelopment()
	// logger.Info("NewDevelopment")

	// logger, _ = zap.NewProduction()
	// logger.Info("NewProduction")

	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)

	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info log", zap.Int("Line", 1))
	logger.Error("Info  Error", zap.Int("Line", 2))
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

func getWriterSync() zapcore.WriteSyncer {

	filePath := "./log/log.txt"

	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		myFile, err := os.Create(filePath)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(myFile)
	}

	file, _ := os.OpenFile(filePath, os.O_RDWR, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	synConsole := zapcore.AddSync(os.Stderr)

	return zapcore.NewMultiWriteSyncer(synConsole, syncFile)
}
