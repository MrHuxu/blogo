package postSvc

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "os"
	"strconv"
)

func (pSvc *PostSvc) ShowSnippets(c *gin.Context) {
	// c.HTML(http.StatusOK, "index.tmpl", gin.H{
	// 	"prd":   "Production" == os.Getenv("ENV"),
	// 	"title": "Blogo",
	// })
	param := c.Param("page")
	page, err := strconv.Atoi(param)
	CheckErr(err)

	result := make(map[string]*Post)
	subTitles := pSvc.Titles[10*page : 10*(page+1)]
	for i := range subTitles {
		result[subTitles[i]] = pSvc.Posts[subTitles[i]]
	}
	c.JSON(http.StatusOK, gin.H{
		"titles": subTitles,
		"posts":  result,
	})
}

func (pSvc *PostSvc) ShowSinglePost(c *gin.Context) {
	title := c.Param("title")
	pSvc.Posts[title].GetTotalContent()
	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"post":  pSvc.Posts[title],
	})
}

func (pSvc *PostSvc) ShowArchives(c *gin.Context) {
	c.JSON(http.StatusOK, pSvc)
}
