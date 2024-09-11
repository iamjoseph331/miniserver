package log

import (
	"context"
	"fmt"

	"github.com/iamjoseph331/miniserver/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	analysisType = "A"
	debugType    = "D"
	requestID    = "X-Kong-Request-ID"
)

var Logger *zap.Logger

func ApplicationLog(ctx context.Context, format string, val ...any) (string, zap.Field, zap.Field) {
	return fmt.Sprintf(format, val...), zap.String("type", debugType), zap.String("request_id", getRequestID(ctx))
}

func AnalysisLog(msg string) (string, zap.Field) {
	return msg, zap.String("type", analysisType)
}

func Setup() {
	cfg := zap.Config{
		Encoding:         "json",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "message",
			LevelKey:     "level",
			EncodeLevel:  zapcore.CapitalLevelEncoder,
			TimeKey:      "time",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	switch config.Conf.Logger.Level {
	case "DEBUG":
		cfg.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	case "INFO":
		cfg.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	case "WARN":
		cfg.Level = zap.NewAtomicLevelAt(zapcore.WarnLevel)
	case "ERROR":
		cfg.Level = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	default:
		cfg.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	}

	var err error
	Logger, err = cfg.Build()
	if err != nil {
		panic(err)
	}
}

func getRequestID(ctx context.Context) string {
	requestID := ctx.Value(requestID)
	if requestID == nil {
		return ""
	}
	return requestID.(string)
}
