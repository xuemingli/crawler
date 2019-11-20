package view

import (
	"learngo/crawler/engine"
	"learngo/crawler/frontend/model"
	common "learngo/crawler/model"
	"os"
	"testing"
)

func TestSearchResult(t *testing.T) {
	view := CreateSearchResultView("template.html")
	out, err := os.Create("template.test.html")
	page := model.SearchResult{}
	page.Hits = 123
	item := engine.Item{
		Url:  "https://album.zhenai.com/u/22576276",
		Type: "zhenai",
		Id:   "22576276",
		Payload: common.Profile{
			Name:       "Eric",
			Age:        37,
			Gender:     "男",
			Height:     182,
			Weight:     100,
			Income:     "12001-20000元",
			Marriage:   "未婚",
			Education:  "大学本科",
			Birthplace: "北京",
		},
	}
	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}
	err = view.Rander(out, page)
	if err != nil {
		panic(err)
	}
}
