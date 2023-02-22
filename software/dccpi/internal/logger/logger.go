package logger

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"dccpi.log",
	}
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC1123)

	return cfg.Build()
}
