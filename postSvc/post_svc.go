package postSvc

import (
	"io/ioutil"
)

type PostSvc struct {
	Posts map[string]*Post // map post title to post entities
}

func New() *PostSvc {
	pSvc := PostSvc{make(map[string]*Post)}

	var files, err = ioutil.ReadDir("./archives")
	CheckErr(err)
	for i := range files {
		p := GetInfosFromName(files[i].Name())
		pSvc.Posts[p.Title] = p
	}

	return &pSvc
}
