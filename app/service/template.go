package service

import (
	"github.com/gin-gonic/gin"
	"html/template"
)

var funcMap = template.FuncMap{
	"add": func(i, j int) int {
		return i + j
	},

	"sub": func(i, j int) int {
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

	"getPicSequence": func(seq int) int {
		result := seq - 13
		if result < 0 {
			result = 81 + result
		}
		return result
	},

	"getPicPosition": func(index int) string {
		if index%2 == 1 {
			return "left"
		}
		return "right"
	},

	"getContentPosition": func(index int) string {
		if index%2 == 1 {
			return "right"
		}
		return "left"
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
