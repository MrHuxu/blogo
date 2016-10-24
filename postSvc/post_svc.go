package postSvc

import (
	"io/ioutil"
	"math"
	"sort"
)

type PostSvc struct {
	MaxPage int              `json:"maxPage"`
	Pages   []int            `json:"pages"`
	Titles  []string         `json:"titles"`
	Posts   map[string]*Post `json:"posts"` // map post titles to post entities
}

func (pSvc PostSvc) Len() int {
	return len(pSvc.Titles)
}

func (pSvc PostSvc) Swap(i, j int) {
	pSvc.Titles[i], pSvc.Titles[j] = pSvc.Titles[j], pSvc.Titles[i]
}

func (pSvc PostSvc) Less(i, j int) bool {
	return pSvc.Posts[pSvc.Titles[i]].Seq > pSvc.Posts[pSvc.Titles[j]].Seq
}

func (pSvc *PostSvc) GeneratePages() {
	i := 0
	for ; i < int(math.Ceil(float64(len(pSvc.Titles)/10.0))); i++ {
		pSvc.Pages = append(pSvc.Pages, i)
	}
	pSvc.MaxPage = i
}

func (pSvc *PostSvc) CachePosts() {
	var files, err = ioutil.ReadDir("./archives")
	CheckErr(err)
	for i := range files {
		p := GetInfosFromName(files[i].Name())
		pSvc.Posts[p.Title] = p
		pSvc.Titles = append(pSvc.Titles, p.Title)
	}
	sort.Sort(pSvc)
}

func New() *PostSvc {
	pSvc := PostSvc{
		Titles: []string{},
		Posts:  make(map[string]*Post),
	}
	pSvc.CachePosts()
	pSvc.GeneratePages()

	return &pSvc
}
