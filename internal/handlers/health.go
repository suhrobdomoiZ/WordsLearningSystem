package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/suhrobdomoiZ/WordsLearningSystem/internal/services"
)

type Health struct {
	service *services.Health
}

func (h *Health) Handler(writer http.ResponseWriter, request *http.Request) {
	status := h.service.CheckServerStatus()
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(status)
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func NewHealth() *Health {
	return &Health{
		service: services.NewHealth(),
	}
}
