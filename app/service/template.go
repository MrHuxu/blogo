package service

import (
	"github.com/gin-gonic/gin"
	"html/template"
)

var funcMap = template.FuncMap{
	"add": func(i int, j int) int {
		return i + j
	},

	"sub": func(i int, j int) int {
		return i - j
	},

	"mul": func(i int, j float32) int {
		return int(float32(i) * j)
	},

	"eq": func(str1, str2 string) bool {
		return str1 == str2
	},

	"not": func(b bool) bool {
		return !b
	},

	"removeYear": func(date string) string {
		return date[:6]
	},
}

func (svc *Service) SetTemplates(server *gin.Engine) error {
	tmpl, err := template.New("").Funcs(funcMap).ParseGlob("./app/templates/*.tmpl")

	if err != nil {
		return err
	}
	server.SetHTMLTemplate(tmpl)
	return nil
}
