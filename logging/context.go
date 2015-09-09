package logging

import "golang.org/x/net/context"

type key string

const loggingKey key = "logging"

func NewContext(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, loggingKey, logger)
}

func FromContext(ctx context.Context) Logger {
	logger, _ := ctx.Value(loggingKey).(Logger)
	return logger
}
