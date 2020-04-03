package main

import (
	"learngo/crawler/engine"
	"learngo/crawler/model"
	"learngo/crawler_distributed/config"
	"learngo/crawler_distributed/rpcsupport"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"
	go serveRpc(host, "test1")
	time.Sleep(time.Second)
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	item := engine.Item{
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
	result := ""
	err = client.Call(config.ItemSaverRpc, item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result: %s; err: %s", result, err)
	}
}
