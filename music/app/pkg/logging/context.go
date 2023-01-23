package logging

import (
	"context"

	"go.uber.org/zap"
)

type ctxLogger struct{}

// ContextWithLogger adds logger to context
func ContextWithLogger(ctx context.Context, l Logger) context.Context {
	return context.WithValue(ctx, ctxLogger{}, l)
}

// LoggerFromContext returns logger from context
func LoggerFromContext(ctx context.Context) *logger {
	if l, ok := ctx.Value(ctxLogger{}).(*logger); ok {
		return l
	}
	return &logger{zap.L()}
}
