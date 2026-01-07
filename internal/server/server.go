package server

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/suhrobdomoiZ/WordsLearningSystem/config"
	"github.com/suhrobdomoiZ/WordsLearningSystem/internal/handlers"
	"github.com/suhrobdomoiZ/WordsLearningSystem/internal/middlewares"
)

type Server struct {
	HTTPServer *http.Server
	logger *slog.Logger
}

func NewServer(port int, logger *slog.Logger) *Server {
	return &Server{
		HTTPServer: &http.Server{
			Addr: fmt.Sprintf(":%d", port),
			ReadHeaderTimeout: config.ReadHeaderTimeout,
		},
		logger: logger,
	}
}

func (s *Server) Start(){
	err := s.HTTPServer.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed){
		s.logger.Error("start server error", slog.Any("error", err))
	}
}

func (s *Server) AddHandlers() {
	health := handlers.NewHealth()
	homepage := handlers.NewHomepage()
	mux := http.NewServeMux()
	mux.HandleFunc("/health/", health.Handler)
	mux.HandleFunc("/", homepage.Handler)

	s.HTTPServer.Handler = mux
}

func (s *Server) AddMidlewares(){
	handler := s.HTTPServer.Handler
	handler = middlewares.Logging(s.logger, handler)
	s.HTTPServer.Handler = handler
}