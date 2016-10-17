package main

import (
	"io/ioutil"
	"strings"
	"time"
)

type Post struct {
	Name    string    `json:"name"`
	Seq     int       `json:"seq"`
	Title   string    `json:"title"`
	Time    time.Time `json:"time"`
	Tags    []string  `json:"tags"`
	Content string    `json:"content"`
}

func (p *Post) GetPartialContent() {

}

func (p *Post) GetTotalContent() {

}

func GetInfosFromName(name string) *Post {
	var infoArr = strings.Split(name, "*")
}

func GetPagedSnippets(page int) []*Post {
	var files, err = ioutil.ReadDir("../archives")
	CheckErr(err)

	var result []*Post
	for i := range files {
		result = append(result, GetInfosFromName(files[i].Name()))
	}
	return result
}

func GetSinglePost(name string) *Post {
	return &Post{}
}
