package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func newZapLogger(logPath string) *zap.Logger {
	// Configure Lumberjack for rolling logs
	writeSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    10,   // Max size in MB before rolling
		MaxBackups: 5,    // Max number of old logs to retain
		MaxAge:     30,   // Max number of days to retain old logs
		Compress:   true, // Compress old logs
		LocalTime:  true,
	})

	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	core := zapcore.NewCore(encoder, writeSyncer, zapcore.InfoLevel)
	return zap.New(core)
}
