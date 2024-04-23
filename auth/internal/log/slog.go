package log

import (
	"context"
	"log/slog"
	"os"
)

type Logger struct {
	*slog.Logger
}

type CtxKey string

const (
	CtxLogger CtxKey = "logger"
)

func WithLogger(ctx context.Context, log Logger) context.Context {
	return context.WithValue(ctx, CtxLogger, log)
}

func GetLogger(ctx context.Context) Logger {
	l, _ := ctx.Value(CtxLogger).(Logger)
	return l
}

func NewLogger(env string) Logger {
	var l *slog.Logger

	switch env {
	case "local":
		l = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case "prod":
		l = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	default:
		panic("unknown env: " + env)
	}

	return Logger{l}
}
