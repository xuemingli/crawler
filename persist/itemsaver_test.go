package persist

import (
	"context"
	"encoding/json"
	"learngo/crawler/engine"
	"learngo/crawler/model"
	"testing"

	"github.com/olivere/elastic/v7"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
		Url:  "https://album.zhenai.com/u/22576276",
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

	// TODO: Try to start up elasticsearch, here using docker run elasticsearch client.
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	const index = "dating_test"
	err = Save(client, expected, index)
	if err != nil {
		panic(err)
	}

	resp, err := client.Get().Index(index).Type(expected.Type).Id(expected.Id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	t.Logf("%s", resp.Source)
	var actual engine.Item
	json.Unmarshal([]byte(resp.Source), &actual)

	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile
	if actual != expected {
		t.Errorf("Got %v; expected %v", actual, expected)
	}
}
