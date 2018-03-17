package service

import (
	"github.com/gin-gonic/gin"
	"html/template"
)

var getYear = func(date string) string {
	return date[8:]
}

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

	"getYear": getYear,

	"removeYear": func(date string) string {
		return date[:6]
	},

	"initYears": func() *[]string {
		return &([]string{})
	},

	"shouldRenderYear": func(existingYears *[]string, date string) bool {
		year := getYear(date)
		for _, existingYear := range *existingYears {
			if existingYear == year {
				return false
			}
		}
		*existingYears = append(*existingYears, year)
		return true
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
