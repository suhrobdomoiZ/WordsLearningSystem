package config

import "errors"

type ConfigError error

var (
	ErrInvalidPort ConfigError = errors.New("invalid port number")
)