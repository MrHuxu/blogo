package service

import (
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Service struct {
	maxPage int
	titles  []string
	tags    []string
	posts   map[string]*post
}

func (svc *Service) postListCanBeAppend(page int) bool {
	return page < svc.maxPage
}

func (svc *Service) paginatedTitles(page int) []string {
	var result []string
	if page >= svc.maxPage {
		result = svc.titles[10*page : len(svc.titles)]
	} else {
		result = svc.titles[10*page : 10*(page+1)]
	}
	return result
}

func (svc *Service) filterByTag(selectedTag string) []string {
	var result []string
	var flag bool
	var yearFlag string

	if selectedTag == "" {
		for _, title := range svc.titles {
			year := strconv.Itoa(svc.posts[title].Date.Year())
			if year != yearFlag {
				result = append(result, year)
				yearFlag = year
			}

			result = append(result, title)
		}
	} else {
		for _, title := range svc.titles {
			flag = false
			for _, tag := range svc.posts[title].Tags {
				flag = tag == selectedTag
				if flag {
					break
				}
			}
			if flag {
				year := strconv.Itoa(svc.posts[title].Date.Year())
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
			p, err := getInfosFromName(file.Name())
			if err != nil {
				return err
			}

			svc.posts[p.Title] = p
			svc.titles = append(svc.titles, p.Title)
		}
	}
	sort.Slice(svc.titles, func(i, j int) bool {
		return svc.posts[svc.titles[i]].Seq > svc.posts[svc.titles[j]].Seq
	})
	return nil
}

func (svc *Service) getMaxPage() {
	svc.maxPage = int(math.Ceil(float64(len(svc.titles) / 10.0)))
}

func New() (*Service, error) {
	service := &Service{
		titles: []string{},
		posts:  make(map[string]*post),
	}

	err := service.cachePosts()
	if err != nil {
		return nil, err
	}

	service.getMaxPage()

	return service, nil
}
