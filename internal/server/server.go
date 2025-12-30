package server

import (
	"fmt"
	"net/http"

	"github.com/suhrobdomoiZ/WordsLearningSystem/internal/handlers"
)

type Server struct {
	HTTPServer *http.Server
}

func NewServer(port int) *Server {
	return &Server{
		HTTPServer: &http.Server{
			Addr: fmt.Sprintf(":%d", port),
		},
	}
}

func (s *Server) Start() error {
	return s.HTTPServer.ListenAndServe()
}

func (s *Server) AddHandlers() {
	health := handlers.NewHealth()
	mux := http.NewServeMux()
	mux.HandleFunc("/health/", health.Handler)

	s.HTTPServer.Handler = mux
}
