package server

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/suhrobdomoiZ/WordsLearningSystem/config"
	"github.com/suhrobdomoiZ/WordsLearningSystem/internal/handlers"
)

type Server struct {
	HTTPServer *http.Server
	logger *slog.Logger
}

func NewServer(port int) *Server {
	return &Server{
		HTTPServer: &http.Server{
			Addr: fmt.Sprintf(":%d", port),
			ReadHeaderTimeout: config.ReadHeaderTimeout,
		},
	}
}

func (s *Server) Start() error {
	return s.HTTPServer.ListenAndServe()
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

}