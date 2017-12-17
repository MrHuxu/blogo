package service

import (
	"github.com/russross/blackfriday"
	"html/template"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type post struct {
	Name     string        `json:"name"`
	Seq      int           `json:"seq"`
	Title    string        `json:"title"`
	Date     time.Time     `json:"date"`
	ShowDate string        `json:"showDate"`
	Tags     []string      `json:"tags"`
	Content  template.HTML `json:"content"`
}

func (p *post) GetFileContent() string {
	data, err := ioutil.ReadFile("./archives/" + p.Name)
	CheckErr(err)
	return string(data)
}

func (p *post) GetPartialContent() {
	bytes := []byte(p.GetFileContent()[0:200])
	p.Content = template.HTML(blackfriday.MarkdownCommon(bytes))
}

func (p *post) GetTotalContent() {
	bytes := []byte(p.GetFileContent())
	p.Content = template.HTML(blackfriday.MarkdownCommon(bytes))
}

func GetInfosFromName(name string) *Post {
	infoArr := strings.Split(name, "*")
	seq, err := strconv.Atoi(infoArr[0])
	CheckErr(err)
	title := infoArr[1]
	date, err := time.Parse("20060102", infoArr[2])
	showDate := date.Format("Jan 02, 2006")
	CheckErr(err)
	tags := strings.Split(strings.Split(infoArr[3], ".")[0], "-")

	return &Post{name, seq, title, date, showDate, tags, ""}
}
