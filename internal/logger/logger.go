package logger

import (
	"go.uber.org/zap"
)

var (
	Logger     *zap.Logger
	HttpLogger *zap.Logger
)

func init() {
	InitLogger()
}

func InitLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err.Error())
	}
	httpLogger, err := zap.NewProduction()
	if err != nil {
		panic(err.Error())
	}
	Logger = logger
	HttpLogger = httpLogger
	return Logger
}

func InitWriterLogger() *zap.Logger {
	Logger = newZapLogger("/tmp/your-loc/file.log")
	HttpLogger = newZapLogger("/tmp/your-loc/file_http.log")
	return Logger
}
