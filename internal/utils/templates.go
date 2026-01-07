package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"time"

	"github.com/suhrobdomoiZ/WordsLearningSystem/internal/assets"
	"github.com/suhrobdomoiZ/WordsLearningSystem/internal/dto"
)

type templateData struct {
	Data any
	User *dto.User
	Year int
}

func GetTemplate(templates ...string) (*template.Template, error) {
	templates = append(
		templates,
		"templates/header.html",
		"templates/footer.html",
		"templates/base.html",
	)

	pageTemplate, err := template.ParseFS(assets.GetTemplatesFS(), templates...)
	if err != nil {
		return nil, fmt.Errorf("error parsing templates: %w", err)
	}

	return pageTemplate, nil
}

func WriteTemplate(
	responseWriter http.ResponseWriter,
	pageTemplate *template.Template,
	data any,
	user *dto.User,
) error {
	pageData := templateData{
		Data: data,
		User: user,
		Year: time.Now().Year(),
	}

	var buffer bytes.Buffer

	err := pageTemplate.Execute(&buffer, pageData)
	if err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}

	responseWriter.Header().Set("Content-Type", "text/html; charset=utf-8")

	_, err = io.Copy(responseWriter, &buffer)
	if err != nil {
		return fmt.Errorf("error writing template: %w", err)
	}

	return nil
}
