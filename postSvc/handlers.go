package postSvc

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

var atPrd = os.Getenv("ENV") == "Production"

func (pSvc *PostSvc) ShowSnippets(c *gin.Context) {
	param := c.Param("page")
	page, err := strconv.Atoi(param)
	CheckErr(err)

	hasPrev, hasNext := pSvc.HasPrevOrNext(page)
	paginatedTitles := pSvc.PaginatedTitles(page)

	paginatedPosts := make(map[string]*Post)
	for i := range paginatedTitles {
		paginatedPosts[paginatedTitles[i]] = pSvc.Posts[paginatedTitles[i]]
		paginatedPosts[paginatedTitles[i]].GetPartialContent()
	}

	c.HTML(http.StatusOK, "layout", gin.H{
		"prd":         atPrd,
		"homePage":    true,
		"pageTitle":   "Life of xhu - Page " + param,
		"pages":       pSvc.Pages,
		"currentPage": page,
		"titles":      paginatedTitles,
		"posts":       paginatedPosts,
		"hasPrev":     hasPrev,
		"hasNext":     hasNext,
	})
}

func (pSvc *PostSvc) ShowSinglePost(c *gin.Context) {
	title := c.Param("title")
	pSvc.Posts[title].GetTotalContent()
	c.HTML(http.StatusOK, "layout", gin.H{
		"prd":       atPrd,
		"postPage":  true,
		"pageTitle": "Life of xhu - " + title,
		"title":     title,
		"post":      pSvc.Posts[title],
	})
}

func (pSvc *PostSvc) ShowArchives(c *gin.Context) {
	c.HTML(http.StatusOK, "layout", gin.H{
		"prd":          atPrd,
		"archivesPage": true,
		"pageTitle":    "Life of xhu - Archives",
		"selectedTag":  c.Query("tag"),
		"titles":       pSvc.FilterByTag(c.Query("tag")),
		"rawData":      pSvc,
	})
}
