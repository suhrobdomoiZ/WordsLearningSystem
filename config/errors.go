package config

import "errors"

type ConfigError error

var (
	ErrInvalidPort              ConfigError = errors.New("invalid port number")
	ErrInvalidLoggerLevel       ConfigError = errors.New("invalid logger level")
	ErrEmptyLoggerLevel         ConfigError = errors.New("empty logger level")
	ErrEmptyLoggerHandlerType   ConfigError = errors.New("empty logger handler type")
	ErrInvalidLoggerHandlerType ConfigError = errors.New("invalid logger handler type")
)
