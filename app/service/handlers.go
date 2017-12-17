package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

var atPrd = os.Getenv("ENV") == "Production"

func handleError(err error) { panic(err) }

func (svc *Service) homeHandler(c *gin.Context) {
	param := c.Param("page")
	page, err := strconv.Atoi(param)
	handleError(err)

	canBeAppend := svc.postListCanBeAppend(page)
	paginatedTitles := svc.paginatedTitles(page)

	paginatedPosts := make(map[string]*post)
	for i := range paginatedTitles {
		paginatedPosts[paginatedTitles[i]] = svc.Posts[paginatedTitles[i]]
		paginatedPosts[paginatedTitles[i]].GetPartialContent()
	}

	c.HTML(http.StatusOK, "layout", gin.H{
		"prd":         atPrd,
		"homePage":    true,
		"pageTitle":   "Life of xhu - Page " + param,
		"currentPage": page,
		"canBeAppend": canBeAppend,
		"titles":      paginatedTitles,
		"posts":       paginatedPosts,
	})
}

func (svc *Service) postHandler(c *gin.Context) {
	title := c.Param("title")
	svc.Posts[title].GetTotalContent()
	c.HTML(http.StatusOK, "layout", gin.H{
		"prd":       atPrd,
		"postPage":  true,
		"pageTitle": "Life of xhu - " + title,
		"title":     title,
		"post":      svc.Posts[title],
	})
}

func (svc *Service) archivesHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "layout", gin.H{
		"prd":          atPrd,
		"archivesPage": true,
		"pageTitle":    "Life of xhu - Archives",
		"selectedTag":  c.Query("tag"),
		"titles":       svc.filterByTag(c.Query("tag")),
		"rawData":      svc,
	})
}
