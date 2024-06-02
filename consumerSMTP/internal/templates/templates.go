package templates

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"

	"consumerSMTP/internal/entities"
)

//go:embed template.gohtml
var templateFS embed.FS

const mimeType = "MIME-version: 1.0;\\nContent-Type: text/html; charset=\\\"UTF-8\\\";\\n\\n"

type Template struct {
}

func New() Template {
	return Template{}
}

func (t Template) Render(data []entities.Currency, subject string) (bytes.Buffer, error) {
	tmpl, err := template.ParseFS(templateFS)
	if err != nil {
		return bytes.Buffer{}, err
	}

	var buffer bytes.Buffer

	buffer.Write([]byte(fmt.Sprintf("Subject: %s\n", subject)))
	buffer.Write([]byte(mimeType))

	err = tmpl.Execute(&buffer, data)
	if err != nil {
		return bytes.Buffer{}, err
	}

	return bytes.Buffer{}, err
}
