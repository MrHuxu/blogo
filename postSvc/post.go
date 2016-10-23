package postSvc

import (
	"github.com/russross/blackfriday"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type Post struct {
	Name    string    `json:"name"`
	Seq     int       `json:"seq"`
	Title   string    `json:"title"`
	Date    time.Time `json:"date"`
	Tags    []string  `json:"tags"`
	Content string    `json:"content"`
}

func (p *Post) GetFileContent() string {
	data, err := ioutil.ReadFile("./archives/" + p.Name)
	CheckErr(err)
	return string(data)
}

func (p *Post) GetPartialContent() {
	bytes := []byte(p.GetFileContent()[0:800])
	p.Content = string(blackfriday.MarkdownCommon(bytes))
}

func (p *Post) GetTotalContent() {
	bytes := []byte(p.GetFileContent())
	p.Content = string(blackfriday.MarkdownCommon(bytes))
}

func GetInfosFromName(name string) *Post {
	infoArr := strings.Split(name, "*")
	seq, err := strconv.Atoi(infoArr[0])
	CheckErr(err)
	title := infoArr[1]
	date, err := time.Parse("20060102", infoArr[2])
	CheckErr(err)
	tags := strings.Split(strings.Split(infoArr[3], ".")[0], "-")

	return &Post{name, seq, title, date, tags, ""}
}

func GetPagedSnippets(page int) []*Post {
	var files, err = ioutil.ReadDir("../archives")
	CheckErr(err)

	var result []*Post
	for i := range files {
		p := GetInfosFromName(files[i].Name())
		p.GetPartialContent()
		result = append(result, p)
	}
	return result
}

func GetSinglePost(name string) *Post {
	p := GetInfosFromName(name)
	p.GetTotalContent()
	return p
}
