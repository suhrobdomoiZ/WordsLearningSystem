package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Port int
}

func getPort() (int, ConfigError) {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return 0, fmt.Errorf("%w: %w", ErrInvalidPort, err)
	}
	return port, nil
}
func LoadConfig() (*Config, ConfigError) {
	port, err := getPort()
	if err != nil {
		return nil, err
	}
	return &Config{
		Port: port,
	}, nil
}
