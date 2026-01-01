package main

import (
	"github.com/suhrobdomoiZ/WordsLearningSystem/internal/server"
	"github.com/suhrobdomoiZ/WordsLearningSystem/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	srv := server.NewServer(cfg.Port)
	srv.AddHandlers()
	err = srv.Start()
	if err != nil {
		panic(err)
	}
}
