package config

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/suhrobdomoiZ/WordsLearningSystem/pkg/logger"
)

const (
	ReadHeaderTimeout time.Duration = 5 * time.Second
	LoggerLevelKey                  = "LOGGER_LEVEL"
	PortKey                         = "PORT"
	LoggerHandlerKey                = "LOGGER_HANDLER_TYPE"
)

type Config struct {
	Port          int
	LoggerLevel   slog.Level
	LoggerHandler logger.HandlerType
}

func getPort() (int, ConfigError) {
	port, err := strconv.Atoi(os.Getenv(PortKey))
	if err != nil {
		return 0, fmt.Errorf("%w: %w", ErrInvalidPort, err)
	}

	return port, nil
}

func getLevel() (slog.Level, ConfigError) {
	level := os.Getenv(LoggerLevelKey)
	switch strings.ToUpper(level) {
	case slog.LevelDebug.String():
		return slog.LevelDebug, nil
	case slog.LevelInfo.String():
		return slog.LevelInfo, nil
	case slog.LevelError.String():
		return slog.LevelError, nil
	case slog.LevelWarn.String():
		return slog.LevelWarn, nil
	}

	if level == "" {
		return slog.LevelInfo, ErrEmptyLoggerLevel
	}

	return slog.LevelInfo, ErrInvalidLoggerLevel
}

func getLoggerHandlerType() (logger.HandlerType, ConfigError) {
	handlerType := os.Getenv(LoggerHandlerKey)
	switch logger.HandlerType(handlerType) {
	case logger.HandlerText:
		return logger.HandlerText, nil
	case logger.HandlerJSON:
		return logger.HandlerJSON, nil
	}

	if handlerType == "" {
		return logger.HandlerType(""), ErrEmptyLoggerHandlerType
	}

	return logger.HandlerType(""), ErrInvalidLoggerHandlerType
}

func LoadConfig() (*Config, ConfigError) {
	port, err := getPort()
	if err != nil {
		return nil, err
	}

	loggerLevel, err := getLevel()
	if err != nil {
		return nil, err
	}

	handlerType, err := getLoggerHandlerType()
	if err != nil {
		return nil, err
	}

	return &Config{
		Port:          port,
		LoggerLevel:   loggerLevel,
		LoggerHandler: handlerType,
	}, nil
}
