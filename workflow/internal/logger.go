package internal

import (
	"go.uber.org/zap"
	"os"
)

func NewLogger() (*zap.Logger, error) {
	var cfg zap.Config
	if os.Getenv("ENV") != "production" {
		cfg = zap.NewDevelopmentConfig()
	} else {
		cfg = zap.NewProductionConfig()
	}

	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}
