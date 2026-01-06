package config

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"time"

	"github.com/suhrobdomoiZ/WordsLearningSystem/pkg/logger"
)

const ReadHeaderTimeout time.Duration = 5 * time.Second
const LoggerLevelKey = "LOGGER_LEVEL"
const PortKey = "PORT"
const LoggerHandlerKey = "LOGGER_HANDLER_TYPE"

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
	switch level {
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
		return slog.LevelInfo, fmt.Errorf("%w", ErrEmptyLoggerLevel)
	}

	return slog.LevelInfo, fmt.Errorf("%w", ErrInvalidLoggerLevel)

}

func getLoggerHandlerType() (logger.HandlerType, error){
	tp := os.Getenv(LoggerHandlerKey)
	switch logger.HandlerType(tp){
	case logger.HandlerText:
		return logger.HandlerText, nil
	case logger.HandlerJSON:
		return logger.HandlerJSON, nil
	}

	if tp == "" {
		return logger.HandlerType(""), fmt.Errorf("%w", ErrEmptyLoggerHandlerType)
	}

	return logger.HandlerType(""), fmt.Errorf("%w", ErrInvalidLoggerHandlerType)
}

func LoadConfig() (*Config, ConfigError) {
	port, err := getPort()
	if err != nil {
		return nil, err
	}
	loggerLevel, errr := getLevel()
	if errr != nil {
		return nil, errr
	}
	tp, errrr := getLoggerHandlerType()
	if errrr != nil{
		return nil, errrr
	}

	return &Config{
		Port:        port,
		LoggerLevel: loggerLevel,
		LoggerHandler: tp,
	}, nil
}
