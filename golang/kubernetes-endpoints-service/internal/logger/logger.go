package logger

import (
	"os"
	"strings"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Log to use application log configuration
var Log *zap.SugaredLogger

// InitializeLogger to initialize the application logger
func InitializeLogger() {

	logLevel := zapcore.InfoLevel
	if strings.ToLower(viper.GetString("log.level")) == "debug" {
		logLevel = zapcore.DebugLevel
	}

	core := zapcore.NewCore(getFileEncoder(), getLogWriter(), logLevel)
	if strings.ToLower(viper.GetString("env")) == "dev" {
		core = zapcore.NewTee(
			zapcore.NewCore(getConsoleEncoder(), zapcore.AddSync(os.Stdout), logLevel),
			zapcore.NewCore(getFileEncoder(), getLogWriter(), logLevel),
		)
	}
	Log = zap.New(core, zap.AddCaller()).Sugar()
}

func getConsoleEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getFileEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	logFile := viper.GetString("log.file")
	logFileMaxSize := viper.GetInt("log.file.maxsize")
	logFileMaxBackups := viper.GetInt("log.file.maxbackups")
	logFileMaxAge := viper.GetInt("log.file.maxage")
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    logFileMaxSize, //megabytes
		MaxBackups: logFileMaxBackups,
		MaxAge:     logFileMaxAge, //days
		Compress:   true,
		LocalTime:  true,
	}
	return zapcore.AddSync(lumberJackLogger)
}
