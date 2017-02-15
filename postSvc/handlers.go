package postSvc

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

var atPrd = os.Getenv("ENV") == "Production"

func (pSvc *PostSvc) ShowHomepage(c *gin.Context) {
	c.HTML(http.StatusOK, "layout", gin.H{
		"prd":         atPrd,
		"homePage":    true,
		"pageTitle":   "Life of xhu",
		"selectedTag": c.Query("tag"),
		"titles":      pSvc.FilterByTag(c.Query("tag")),
		"posts":       pSvc.Posts,
		"rawData":     pSvc,
	})
}

func (pSvc *PostSvc) ShowSinglePost(c *gin.Context) {
	title := c.Param("title")
	pSvc.Posts[title].GetTotalContent()
	c.HTML(http.StatusOK, "layout", gin.H{
		"prd":         atPrd,
		"postPage":    true,
		"pageTitle":   "Life of xhu - " + title,
		"title":       title,
		"post":        pSvc.Posts[title],
		"selectedTag": c.Query("tag"),
		"titles":      pSvc.FilterByTag(c.Query("tag")),
		"posts":       pSvc.Posts,
		"rawData":     pSvc,
	})
}
