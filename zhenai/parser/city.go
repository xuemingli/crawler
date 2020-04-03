package parser

import (
	"learngo/crawler/engine"
	"regexp"
)

//const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

var cityRe2 = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a></th></tr> <tr><td width="180"><span class="grayL">性别：</span>([^<]+)</td>`)

var cityNextPageRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)

func ParseCity(contents []byte) engine.ParseResult {

	submatch := cityRe2.FindAllSubmatch(contents, -1)
	//for _, m := range submatch {
	//	fmt.Printf(">>>%s\n", string(m[1]))
	//}

	result := engine.ParseResult{}
	for _, m := range submatch {
		name := string(m[2])
		m[3] = m[3][:len(m[3])-3]
		gender := string(m[3])
		url := string(m[1])
		//result.Items = append(result.Items, "User "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: NewProfileParser(name, url, gender),
		})
	}

	submatch1 := cityNextPageRe.FindAllSubmatch(contents, -1)
	for _, m := range submatch1 {
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: engine.NewFuncParser(ParseCity, "ParseCity"),
		})
	}
	return result
}
