package postSvc

import (
	"github.com/russross/blackfriday"
	"html/template"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type Post struct {
	Name    string        `json:"name"`
	Seq     int           `json:"seq"`
	Title   string        `json:"title"`
	Date    time.Time     `json:"date"`
	Tags    []string      `json:"tags"`
	Content template.HTML `json:"content"`
}

func (p *Post) GetFileContent() string {
	data, err := ioutil.ReadFile("./archives/" + p.Name)
	CheckErr(err)
	return string(data)
}

func (p *Post) GetPartialContent() {
	bytes := []byte(p.GetFileContent()[0:200])
	p.Content = template.HTML(blackfriday.MarkdownCommon(bytes))
}

func (p *Post) GetTotalContent() {
	bytes := []byte(p.GetFileContent())
	p.Content = template.HTML(blackfriday.MarkdownCommon(bytes))
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
