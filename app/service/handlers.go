package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

var atPrd = os.Getenv("ENV") == "Production"

func (svc *Service) homeHandler(c *gin.Context) {
	paginatedTitles := svc.paginatedTitles(0)

	paginatedPosts := make(map[string]*post)
	for i := range paginatedTitles {
		paginatedPosts[paginatedTitles[i]] = svc.posts[paginatedTitles[i]]
		paginatedPosts[paginatedTitles[i]].getPartialContent()
	}

	c.HTML(http.StatusOK, "layout", gin.H{
		"prd":        atPrd,
		"homePage":   true,
		"pageTitle":  "Life of xhu - Home",
		"maxPostSeq": svc.posts[svc.titles[0]].Seq,
		"titles":     paginatedTitles,
		"posts":      paginatedPosts,
	})
}

func (svc *Service) pageHandler(c *gin.Context) {
	param := c.Param("page")
	page, err := strconv.Atoi(param)
	if err != nil {
		page = 1
	}

	canBeAppend := svc.postListCanBeAppend(page)
	paginatedTitles := svc.paginatedTitles(page)

	paginatedPosts := make(map[string]*post)
	for i := range paginatedTitles {
		paginatedPosts[paginatedTitles[i]] = svc.posts[paginatedTitles[i]]
		paginatedPosts[paginatedTitles[i]].getPartialContent()
	}

	c.JSON(http.StatusOK, gin.H{
		"currentPage": page,
		"canBeAppend": canBeAppend,
		"titles":      paginatedTitles,
		"posts":       paginatedPosts,
	})
}

func (svc *Service) postHandler(c *gin.Context) {
	title := c.Param("title")
	svc.posts[title].getTotalContent()
	c.HTML(http.StatusOK, "layout", gin.H{
		"prd":       atPrd,
		"postPage":  true,
		"pageTitle": "Life of xhu - " + title,
		"title":     title,
		"post":      svc.posts[title],
	})
}

func (svc *Service) archivesHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "layout", gin.H{
		"prd":          atPrd,
		"archivesPage": true,
		"pageTitle":    "Life of xhu - Archives",
		"selectedTag":  c.Query("tag"),
		"titles":       svc.filterByTag(c.Query("tag")),
		"posts":        svc.posts,
	})
}
