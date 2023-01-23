package templates

import (
	"embed"
	"html/template"
)

//go:embed *
var templatesFS embed.FS

func GetIndexTemplage() (*template.Template, error) {
	tmpl, err := template.New("index.tmpl").ParseFS(templatesFS, "index.tmpl")
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}
