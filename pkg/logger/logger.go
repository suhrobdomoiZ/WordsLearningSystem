package logger

import (
	"log/slog"
	"os"
)
type HandlerType string

const HandlerText HandlerType = "Text"
const HandlerJSON HandlerType = "JSON"

func NewLogger(tp HandlerType, lv slog.Level) *slog.Logger {
	if tp == HandlerText{
		handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: lv,
		})
		return slog.New(handler)

	}else{
		handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: lv,
		})
		return slog.New(handler)
	}

}