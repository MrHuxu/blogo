package postSvc

import (
	"io/ioutil"
	"math"
	"sort"
)

type PostSvc struct {
	MaxPage  int              `json:"maxPage"`
	Pages    []int            `json:"pages"`
	Titles   []string         `json:"titles"`
	Tags     []string         `json:"tags"`
	TagTimes map[string]int   `json:"tagTimes"`
	Posts    map[string]*Post `json:"posts"` // map post titles to post entities
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
	for ; i <= int(math.Ceil(float64(len(pSvc.Titles)/10.0))); i++ {
		pSvc.Pages = append(pSvc.Pages, i)
	}
	pSvc.MaxPage = i - 1
}

func (pSvc *PostSvc) CacheTags(tags []string) {
	for _, tag := range tags {
		_, ok := pSvc.TagTimes[tag]
		pSvc.TagTimes[tag] += 1
		if !ok {
			pSvc.Tags = append([]string{tag}, pSvc.Tags...)
		}
	}
}

func (pSvc *PostSvc) FilterByTag(selectedTag string) []string {
	var result []string
	var flag bool

	if selectedTag == "" {
		result = pSvc.Titles
	} else {
		for _, title := range pSvc.Titles {
			flag = false
			for _, tag := range pSvc.Posts[title].Tags {
				flag = tag == selectedTag
				if flag {
					break
				}
			}
			if flag {
				result = append(result, title)
			}
		}
	}
	return result
}

func (pSvc *PostSvc) CachePosts() {
	var files, err = ioutil.ReadDir("./archives")
	CheckErr(err)
	for i := range files {
		p := GetInfosFromName(files[i].Name())
		pSvc.Posts[p.Title] = p
		pSvc.CacheTags(p.Tags)
		pSvc.Titles = append(pSvc.Titles, p.Title)
	}
	sort.Sort(pSvc)
}

func New() *PostSvc {
	pSvc := PostSvc{
		Titles:   []string{},
		TagTimes: make(map[string]int),
		Posts:    make(map[string]*Post),
	}
	pSvc.CachePosts()
	pSvc.GeneratePages()

	return &pSvc
}
