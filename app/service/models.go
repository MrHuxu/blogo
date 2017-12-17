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

func (p *post) getFileContent() (string, error) {
	data, err := ioutil.ReadFile("./archives/" + p.Name)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (p *post) getPartialContent() {
	fullContent, _ := p.getFileContent()
	if len([]byte(fullContent)) < 201 {
		p.Content = template.HTML(blackfriday.MarkdownCommon([]byte{}))
	} else {
		p.Content = template.HTML(blackfriday.MarkdownCommon([]byte(fullContent[0:200])))
	}
}

func (p *post) getTotalContent() {
	fullContent, _ := p.getFileContent()
	p.Content = template.HTML(blackfriday.MarkdownCommon([]byte(fullContent)))
}

func getInfosFromName(name string) (*post, error) {
	infoArr := strings.Split(name, "*")
	seq, err := strconv.Atoi(infoArr[0])
	if err != nil {
		return nil, err
	}

	title := infoArr[1]
	date, err := time.Parse("20060102", infoArr[2])
	if err != nil {
		return nil, err
	}

	showDate := date.Format("Jan 02, 2006")
	tags := strings.Split(strings.Split(infoArr[3], ".")[0], "-")

	return &post{name, seq, title, date, showDate, tags, ""}, nil
}
