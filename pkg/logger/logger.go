package logger

import (
	"log/slog"
	"os"
	"strings"
)

type LevelLogger uint8

const (
	LevelError LevelLogger = 1 << iota
	LevelWarn
	LevelInfo
	LevelDebug
)

func (level LevelLogger) String() string {
	switch level {
	case LevelError:
		return "error"
	case LevelWarn:
		return "warn"
	case LevelInfo:
		return "info"
	case LevelDebug:
		return "debug"
	default:
		return "info"
	}
}

type Logger struct {
	logger *slog.Logger
}

func New(level string) *Logger {
	var l slog.Level

	switch strings.ToLower(level) {
	case "error":
		l = slog.LevelError
	case "warn":
		l = slog.LevelWarn
	case "info":
		l = slog.LevelInfo
	case "debug":
		l = slog.LevelDebug
	default:
		l = slog.LevelInfo
	}
	handlerOpts := &slog.HandlerOptions{
		Level: l,
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, handlerOpts))
	slog.SetDefault(logger)
	return &Logger{
		logger: logger,
	}
}
