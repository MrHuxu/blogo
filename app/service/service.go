package service

import (
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Service struct {
	MaxPage int
	Titles  []string
	Tags    []string
	Posts   map[string]*post
}

func (svc *Service) postListCanBeAppend(page int) bool {
	return page < svc.MaxPage
}

func (svc *Service) paginatedTitles(page int) []string {
	var result []string
	if page >= svc.MaxPage {
		result = svc.Titles[10*page : len(svc.Titles)]
	} else {
		result = svc.Titles[10*page : 10*(page+1)]
	}
	return result
}

func (svc *Service) filterByTag(selectedTag string) []string {
	var result []string
	var flag bool
	var yearFlag string

	if selectedTag == "" {
		for _, title := range svc.Titles {
			year := strconv.Itoa(svc.Posts[title].Date.Year())
			if year != yearFlag {
				result = append(result, year)
				yearFlag = year
			}

			result = append(result, title)
		}
	} else {
		for _, title := range svc.Titles {
			flag = false
			for _, tag := range svc.Posts[title].Tags {
				flag = tag == selectedTag
				if flag {
					break
				}
			}
			if flag {
				year := strconv.Itoa(svc.Posts[title].Date.Year())
				if year != yearFlag {
					result = append(result, year)
					yearFlag = year
				}

				result = append(result, title)
			}
		}
	}
	return result
}

func (svc *Service) cachePosts() error {
	files, err := ioutil.ReadDir("./archives")
	if err != nil {
		return err
	}

	for _, file := range files {
		if !strings.HasPrefix(file.Name(), "WIP:") {
			p := GetInfosFromName(file.Name())
			svc.Posts[p.Title] = p
			svc.Titles = append(svc.Titles, p.Title)
		}
	}
	sort.Slice(svc.Titles, func(i, j int) bool {
		return svc.Posts[svc.Titles[i]].Seq > svc.Posts[svc.Titles[j]].Seq
	})
	return nil
}

func (svc *Service) getMaxPage() {
	svc.MaxPage = int(math.Ceil(float64(len(svc.Titles) / 10.0)))
}

func New() (*Service, error) {
	service := &Service{
		Titles: []string{},
		Posts:  make(map[string]*post),
	}

	err := service.cachePosts()
	if err != nil {
		return nil, err
	}

	service.getMaxPage()

	return service, nil
}
