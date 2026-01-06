package main

import (
	"github.com/suhrobdomoiZ/WordsLearningSystem/config"
	"github.com/suhrobdomoiZ/WordsLearningSystem/internal/server"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	srv := server.NewServer(cfg.Port)
	srv.AddHandlers()
	srv.AddMidlewares()

	err = srv.Start()
	if err != nil {
		panic(err)
	}
}
