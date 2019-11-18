package main

import (
	"learngo/crawler/engine"
	"learngo/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "https://www.huazhengcaiwu.com/city/",
		ParserFunc: parser.ParseCityList,
	})
}
