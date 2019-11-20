package parser

import (
	"io/ioutil"
	"learngo/crawler/engine"
	"learngo/crawler/model"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	//profile_test_data.html文件内容是https://album.zhenai.com/u/22576276网页源代码
	if err != nil {
		panic(err)
	}
	results := ParseProfile(contents, "http://album.zhenai.com/u/22576276", "Eric", "男")
	if len(results.Items) != 1 {
		t.Errorf("items should contain 1 element; but was %v", results.Items)
	}
	actual := results.Items[0]
	expected := engine.Item{
		Url:  "http://album.zhenai.com/u/22576276",
		Type: "zhenai",
		Id:   "22576276",
		Payload: model.Profile{
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
	if actual != expected {
		t.Errorf("expected %v; but was %v", expected, actual)
	}
}
