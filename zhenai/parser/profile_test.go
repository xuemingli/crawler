package parser

import (
	"io/ioutil"
	"learngo/crawler/model"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}
	results := ParseProfile(contents, "小明")
	if len(results.Items) != 1 {
		t.Errorf("items should contain 1 element; but was %v", results.Items)
	}
	profile := results.Items[0].(model.Profile)
	expected := model.Profile{
		Name:       "小明",
		Age:        0,
		Height:     "163cm",
		Income:     "3000元以下",
		Marriage:   "离异",
		Education:  "高中及以下",
		Birthplace: "大兴安岭",
	}
	if profile != expected {
		t.Errorf("expected %v; but was %v", expected, profile)
	}
}