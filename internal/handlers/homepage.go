package handlers

import (
	"html/template"
	"net/http"

	"github.com/suhrobdomoiZ/WordsLearningSystem/internal/assets"
	"github.com/suhrobdomoiZ/WordsLearningSystem/internal/services"
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
	pageTemplate, err := template.ParseFS(
		assets.GetTemplatesFS(),
		"templates/homepage.html",
		"templates/header.html",
		"templates/base.html",
		"templates/footer.html",
	)
	if err != nil {
		http.Error(
			writer,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)

		return
	}

	data := h.service.GetBuildData()

	err = pageTemplate.Execute(writer, data)
	if err != nil {
		http.Error(
			writer,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
	}
}
