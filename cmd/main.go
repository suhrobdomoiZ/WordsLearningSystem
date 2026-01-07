package main

import (
	"github.com/suhrobdomoiZ/WordsLearningSystem/config"
	"github.com/suhrobdomoiZ/WordsLearningSystem/internal/server"
	"github.com/suhrobdomoiZ/WordsLearningSystem/pkg/logger"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	logger := logger.NewLogger(cfg.LoggerHandler, cfg.LoggerLevel)
	srv := server.NewServer(cfg.Port, logger)

	srv.AddHandlers()
	srv.AddMidlewares()

	srv.Start()
}
