package logger

import "go.uber.org/zap"

func NewLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"../storage/dccpi.log",
	}

	return cfg.Build()
}
