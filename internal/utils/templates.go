package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"time"

	"github.com/suhrobdomoiZ/WordsLearningSystem/internal/assets"
	"github.com/suhrobdomoiZ/WordsLearningSystem/internal/dto"
)

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

func WriteTemplate(responseWriter http.ResponseWriter, pageTemplate *template.Template, data any, user *dto.User) error {
	marshalledData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error marshalling data: %w", err)
	}

	pageData := make(map[string]any)
	err = json.Unmarshal(marshalledData, &pageData)
	if err != nil {
		return fmt.Errorf("error unmarshalling data: %w", err)
	}

	pageData["User"] = user
	pageData["Year"] = time.Now().Year()

	var buffer bytes.Buffer
	err = pageTemplate.Execute(&buffer, data)
	if err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}

	responseWriter.Header().Set("Content-Type", "text/html; charset=utf-8")
	n, err := io.Copy(responseWriter, &buffer)
	if err != nil {
		return fmt.Errorf("error writing template: %w", err)
	}

	if n != int64(buffer.Len()) {
		return fmt.Errorf(
			"error writing template: expected %d bytes, wrote %d bytes",
			buffer.Len(),
			n,
		)
	}

	return nil
}
