package main

import (
	"github.com/MrHuxu/blogo/postSvc"
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

	"not": func(b bool) bool {
		return !b
	},

	"removeYear": func(date string) string {
		return date[:6]
	},
}

func main() {
	server := gin.Default()
	if tmpl, err := template.New("").Funcs(funcMap).ParseGlob("templates/*.tmpl"); err == nil {
		server.SetHTMLTemplate(tmpl)
	} else {
		panic(err)
	}
	server.Static("/assets", "./assets")

	ps := postSvc.New()
	ps.RegisterRoutes(server)

	const port = "13109"
	server.Run(":" + port)
}
