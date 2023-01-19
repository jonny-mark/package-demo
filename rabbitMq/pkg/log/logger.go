package log

import (
	"context"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 全局log
var logger Logger
var zl *zap.Logger

// Fields Type to pass when we want to call WithFields for structured logging
type Fields map[string]interface{}

type Logger interface {
	Info(args ...interface{})
	Infof(format string, args ...interface{})

	Warn(args ...interface{})
	Warnf(format string, args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})

	WithFields(keyValues Fields) Logger
}

func GetLogger() Logger {
	return logger
}

// Info log
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Warn log
func Warn(args ...interface{}) {
	logger.Warn(args...)
}

// Error log
func Error(args ...interface{}) {
	logger.Error(args...)
}

// Infof logger
func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

// Warnf logger
func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

// Errorf logger
func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func WithFields(keyValues Fields) Logger {
	return logger.WithFields(keyValues)
}

func WithContext(ctx context.Context) Logger {
	//return zap logger

	if span := trace.SpanFromContext(ctx); span != nil {
		logger := spanLogger{span: span, logger: zl}

		spanCtx := span.SpanContext()
		logger.spanFields = []zapcore.Field{
			zap.String("trace_id", spanCtx.TraceID().String()),
			zap.String("span_id", spanCtx.SpanID().String()),
		}

		return logger
	}
	return logger
}
