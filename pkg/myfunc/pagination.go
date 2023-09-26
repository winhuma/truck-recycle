package myfunc

import (
	"math"
	"reflect"
	"truck-unii-app/pkg/myresponse"
)

func SetPagination(msg string, rawData interface{}, getPage int, getOffset int) interface{} {
	var res myresponse.PaginationResponse
	var sizePage int
	var page int
	var totalPage int
	var nextPage interface{}
	var allItem int
	var NowShow int
	var data []interface{}
	rv := reflect.ValueOf(rawData)
	if rv.Kind() == reflect.Slice {
		for i := 0; i < rv.Len(); i++ {
			data = append(data, rv.Index(i).Interface())
		}
	}

	allItem = len(data)
	sizePage = getOffset
	page = getPage

	if getOffset == 0 {
		sizePage = 10
	}

	if getPage == 0 {
		page = 1
	}

	divPage := float64(len(data)) / float64(sizePage)
	totalPage = int(math.Ceil(divPage))
	start := (page - 1) * sizePage

	if page > totalPage || allItem == 0 {
		res.Message = msg
		res.Offset = sizePage
		res.Data = []string{}
		res.Page = 1
		res.AllItem = 0
		res.NowShowItem = 0
		res.TotalPage = 1
		return res
	}

	if (page + 1) <= totalPage {
		nextPage = page + 1
	} else {
		nextPage = nil
	}

	data = data[start:]
	if len(data) > sizePage {
		data = data[:sizePage]
	}
	NowShow = len(data)

	res.Message = msg
	res.Data = data
	res.Offset = sizePage
	res.Page = page
	res.NextPage = nextPage
	res.TotalPage = totalPage
	res.NowShowItem = NowShow
	res.AllItem = allItem

	return res
}
