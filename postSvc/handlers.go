package postSvc

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "os"
)

func (pSvc *PostSvc) ShowSnippets(c *gin.Context) {
	// c.HTML(http.StatusOK, "index.tmpl", gin.H{
	// 	"prd":   "Production" == os.Getenv("ENV"),
	// 	"title": "Blogo",
	// })
	for i := range pSvc.Posts {
		pSvc.Posts[i].GetPartialContent()
	}
	c.JSON(http.StatusOK, pSvc.Posts)
}

func (pSvc *PostSvc) ShowSinglePost(c *gin.Context) {
	for i := range pSvc.Posts {
		pSvc.Posts[i].GetTotalContent()
	}
	c.JSON(http.StatusOK, pSvc.Posts)
}

func (pSvc *PostSvc) ShowArchives(c *gin.Context) {
	c.JSON(http.StatusOK, pSvc.Posts)
}
