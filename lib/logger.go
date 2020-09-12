package lib

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger structure
type Logger struct {
	Zap *zap.SugaredLogger
}

// NewLogger sets up logger
func NewLogger(env Env) Logger {

	config := zap.NewDevelopmentConfig()

	if env.Environment == "development" {
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	if env.Environment == "production" && env.LogOutput != "" {
		config.OutputPaths = []string{env.LogOutput}
	}

	logger, _ := config.Build()

	sugar := logger.Sugar()

	return Logger{
		Zap: sugar,
	}

}
