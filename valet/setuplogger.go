package valet

import (
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func init() {
	var err error
	logger, err = buildLogger()
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize logger: %v", err))
	}
	defer logger.Sync() // flushes buffer, if any
}

func buildLogger() (*zap.Logger, error) {
	config := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "severity",
		NameKey:        "logger",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.MillisDurationEncoder,
	}

	logLevelEnv := strings.ToUpper(os.Getenv("LOG_LEVEL"))
	var level zapcore.Level
	switch logLevelEnv {
	case "DEBUG":
		level = zapcore.DebugLevel
	case "INFO":
		level = zapcore.InfoLevel
	case "WARN":
		level = zapcore.WarnLevel
	case "ERROR":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}

	jsonEncoder := zapcore.NewJSONEncoder(config)
	stderr := zapcore.Lock(os.Stderr)
	core := zapcore.NewCore(jsonEncoder, stderr, level)

	return zap.New(core, zap.AddCaller()), nil
}
