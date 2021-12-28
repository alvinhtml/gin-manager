package request

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Page int `json:"page" form:"page"`
	Size int `json:"size" form:"size"`
}

type UnformattedPageQuery struct {
	Pagination
	Filter []string `json:"filter" form:"filter"`
	Search []string `json:"search" form:"search"`
	Sort   []string `json:"sort" form:"sort"`
}

// Paging common input parameter structure
type PageQuery struct {
	Pagination
	Filter map[string]string `json:"filter" form:"filter"`
	Search map[string]string `json:"search" form:"search"`
	Sort   map[string]string `json:"sort" form:"sort"`
}

// Find by id structure
type ById struct {
	Id uint `json:"id" form:"id" uri:"id"`
}

type byIds struct {
	Ids []uint `json:"ids" form:"ids"`
}

// FormatSearch format search parameter
func FormatSearch(unformattedPageQuery UnformattedPageQuery) (err error, searchs map[string]string) {
	if len(unformattedPageQuery.Search) > 0 {
		searchs = make(map[string]string)
		for _, v := range unformattedPageQuery.Search {
			s := strings.Split(v, ":")

			if len(s) != 2 {
				err = errors.New("无法解析 search 参数")
				return
			}

			searchs[s[0]] = s[1]
		}
	}
	return err, searchs
}

// FormatFilter format filter parameter
func FormatFilter(unformattedPageQuery UnformattedPageQuery) (err error, filter map[string]string) {

	if len(unformattedPageQuery.Filter) > 0 {
		filter = make(map[string]string)
		for _, v := range unformattedPageQuery.Filter {
			s := strings.Split(v, ":")

			if len(s) != 2 {
				err = errors.New("无法解析 filter 参数")
				return
			}

			filter[s[0]] = s[1]
		}
	}

	fmt.Println(filter)
	return err, filter
}

// FormatSort format sort parameter
func FormatSort(unformattedPageQuery UnformattedPageQuery) (err error, sort map[string]string) {
	if len(unformattedPageQuery.Sort) > 0 {
		sort = make(map[string]string)
		for _, v := range unformattedPageQuery.Sort {
			order := v[:1]
			if order == "-" {
				sort[v[1:]] = "desc"
			} else if order == "+" {
				sort[v[1:]] = "asc"
			} else {
				sort[v] = "asc"
			}
		}
	}
	return err, sort
}

func (pageQuery *PageQuery) BindQuery(c *gin.Context) (err error) {
	var unformattedPageQuery UnformattedPageQuery
	c.ShouldBindQuery(&unformattedPageQuery)

	fmt.Println(unformattedPageQuery)

	pageQuery.Page = unformattedPageQuery.Page
	pageQuery.Size = unformattedPageQuery.Size

	err, pageQuery.Filter = FormatFilter(unformattedPageQuery)
	err, pageQuery.Search = FormatSearch(unformattedPageQuery)
	err, pageQuery.Sort = FormatSort(unformattedPageQuery)

	return err
}
