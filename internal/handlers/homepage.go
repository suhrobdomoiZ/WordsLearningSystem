package handlers

import (
	"net/http"

	"github.com/suhrobdomoiZ/WordsLearningSystem/internal/services"
	"github.com/suhrobdomoiZ/WordsLearningSystem/internal/utils"
)

type Homepage struct {
	service *services.Homepage
}

func NewHomepage() *Homepage {
	return &Homepage{
		service: services.NewHomepage(),
	}
}

func (h *Homepage) Handler(writer http.ResponseWriter, request *http.Request) {
	pageTemplate, err := utils.GetTemplate("templates/homepage.html")
	if err != nil {
		http.Error(
			writer,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)

		return
	}

	data := h.service.GetBuildData()

	err = utils.WriteTemplate(writer, request, pageTemplate, data)
	if err != nil {
		http.Error(
			writer,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
	}
}
