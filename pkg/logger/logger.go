package logger

import (
	"log/slog"
	"os"
)

type HandlerType string

const (
	HandlerText HandlerType = "Text"
	HandlerJSON HandlerType = "JSON"
)

func NewLogger(handlerType HandlerType, level slog.Level) *slog.Logger {
	var handler slog.Handler

	options := &slog.HandlerOptions{
		Level: level,
	}
	if handlerType == HandlerText {
		handler = slog.NewTextHandler(os.Stdout, options)
	} else {
		handler = slog.NewJSONHandler(os.Stdout, options)
	}

	return slog.New(handler)
}
