package parser

import (
	"learngo/crawler/engine"
	"regexp"
)

var urlprefix string = "http://www.zhenai.com/zhenghun"

//<a href="http://album.zhenai.com/u/1615511201" target="_blank">皮皮虾</a>
var cityListRe = regexp.MustCompile(`<a href="(/[a-z]+/)" >([^<]+)</a>`)

func ParseCityList(contents []byte) engine.ParseResult {
	submatch := cityListRe.FindAllSubmatch(contents, -1)
	//for _, m := range submatch {
	//	fmt.Printf("------%s\n", m)
	//}
	submatch = submatch[6:]
	submatch = submatch[:len(submatch)-6]
	var w string
	result := engine.ParseResult{}
	//limit := 2
	for _, m := range submatch {
		w = urlprefix
		w += string(m[1])
		//result.Items = append(result.Items, "City "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{Url: w, Parser: engine.NewFuncParser(ParseCity, "ParseCity")})
		//fmt.Printf("city: %s, url: %s\n", m[2], w)
		//limit--
		//if limit == 0 {
		//	break
		//}
	}
	return result
}
