package templates

import (
	"embed"
	"html/template"
)

//go:embed index.tmpl
var indexFS embed.FS

func GetIndexTemplage() (*template.Template, error) {
	return nil, nil
}
