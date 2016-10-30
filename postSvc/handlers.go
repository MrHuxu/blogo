package postSvc

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "os"
	"strconv"
)

func (pSvc *PostSvc) ShowSnippets(c *gin.Context) {
	result := make(map[string]*Post)
	var subTitles []string

	param := c.Param("page")
	page, err := strconv.Atoi(param)
	CheckErr(err)

	if page >= pSvc.MaxPage {
		subTitles = pSvc.Titles[10*page : len(pSvc.Titles)]
	} else {
		subTitles = pSvc.Titles[10*page : 10*(page+1)]
	}

	for i := range subTitles {
		result[subTitles[i]] = pSvc.Posts[subTitles[i]]
		result[subTitles[i]].GetPartialContent()
	}

	c.HTML(http.StatusOK, "layout", gin.H{
		"homePage":  true,
		"pageTitle": "Life of xhu - Page " + param,
		"pages":     pSvc.Pages,
		"titles":    subTitles,
		"posts":     result,
	})
}

func (pSvc *PostSvc) ShowSinglePost(c *gin.Context) {
	title := c.Param("title")
	pSvc.Posts[title].GetTotalContent()
	c.HTML(http.StatusOK, "layout", gin.H{
		"postPage":  true,
		"pageTitle": "Life of xhu - " + title,
		"title":     title,
		"post":      pSvc.Posts[title],
	})
}

func (pSvc *PostSvc) ShowArchives(c *gin.Context) {
	c.HTML(http.StatusOK, "layout", gin.H{
		"archivesPage": true,
		"pageTitle":    "Life of xhu - Archive",
		"titles":       pSvc.FilterByTag(c.Query("tag")),
		"rawData":      pSvc,
	})
}
