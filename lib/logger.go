package lib

import (
	"context"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	gormlogger "gorm.io/gorm/logger"
)

// Logger structure
type Logger struct {
	*zap.SugaredLogger
}

type GinLogger struct {
	*Logger
}

type FxLogger struct {
	*Logger
}

type GormLogger struct {
	*Logger
	gormlogger.Config
}

var (
	globalLogger *Logger
	zapLogger    *zap.Logger
)

// GetLogger get the logger
func GetLogger() Logger {
	if globalLogger == nil {
		logger := newLogger()
		globalLogger = &logger
	}
	return *globalLogger
}

// GetGinLogger get the gin logger
func (l Logger) GetGinLogger() GinLogger {
	logger := zapLogger.WithOptions(
		zap.WithCaller(false),
	)
	return GinLogger{
		Logger: newSugaredLogger(logger),
	}
}

// GetFxLogger get the fx logger
func (l Logger) GetFxLogger() FxLogger {
	logger := zapLogger.WithOptions(
		zap.WithCaller(false),
	)

	return FxLogger{
		Logger: newSugaredLogger(logger),
	}
}

// GetGormLogger gets the gorm framework logger
func (l Logger) GetGormLogger() *GormLogger {
	logger := zapLogger.WithOptions(
		zap.AddCaller(),
		zap.AddCallerSkip(3),
	)

	return &GormLogger{
		Logger: newSugaredLogger(logger),
		Config: gormlogger.Config{
			LogLevel: gormlogger.Info,
		},
	}
}

func newSugaredLogger(logger *zap.Logger) *Logger {
	return &Logger{
		SugaredLogger: logger.Sugar(),
	}
}

// newLogger sets up logger
func newLogger() Logger {

	config := zap.NewDevelopmentConfig()
	env := os.Getenv("ENV")
	logOutput := os.Getenv("LOG_OUTPUT")

	if env == "development" {
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	if env == "production" && logOutput != "" {
		config.OutputPaths = []string{logOutput}
	}

	zapLogger, _ = config.Build()
	logger := newSugaredLogger(zapLogger)

	return *logger
}

// Write interface implementation for gin-framework
func (l GinLogger) Write(p []byte) (n int, err error) {
	l.Info(string(p))
	return len(p), nil
}

// Printf prits go-fx logs
func (l FxLogger) Printf(str string, args ...interface{}) {
	l.Infof(str, args)
}

// GORM Framework Logger Interface Implementations
// ---- START ----

// LogMode set log mode
func (l *GormLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

// Info prints info
func (l GormLogger) Info(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel >= gormlogger.Info {
		l.Debugf(str, args...)
	}
}

// Warn prints warn messages
func (l GormLogger) Warn(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel >= gormlogger.Warn {
		l.Warnf(str, args...)
	}

}

// Error prints error messages
func (l GormLogger) Error(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel >= gormlogger.Error {
		l.Errorf(str, args...)
	}
}

// Trace prints trace messages
func (l GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= 0 {
		return
	}
	elapsed := time.Since(begin)
	if l.LogLevel >= gormlogger.Info {
		sql, rows := fc()
		l.Debug("[", elapsed.Milliseconds(), " ms, ", rows, " rows] ", "sql -> ", sql)
		return
	}

	if l.LogLevel >= gormlogger.Warn {
		sql, rows := fc()
		l.SugaredLogger.Warn("[", elapsed.Milliseconds(), " ms, ", rows, " rows] ", "sql -> ", sql)
		return
	}

	if l.LogLevel >= gormlogger.Error {
		sql, rows := fc()
		l.SugaredLogger.Error("[", elapsed.Milliseconds(), " ms, ", rows, " rows] ", "sql -> ", sql)
		return
	}
}
