package templates

import (
	"bytes"
	"embed"
	"html/template"

	"consumerSMTP/internal/entities"
)

//go:embed template.gohtml
var templateFS embed.FS

type Template struct {
}

func New() Template {
	return Template{}
}

func (t Template) Render(data []entities.Currency) (bytes.Buffer, error) {
	tmpl, err := template.ParseFS(templateFS)
	if err != nil {
		return bytes.Buffer{}, err
	}

	var buffer bytes.Buffer

	err = tmpl.Execute(&buffer, data)
	if err != nil {
		return bytes.Buffer{}, err
	}

	return bytes.Buffer{}, err
}
