package parser

import (
	"learngo/crawler/engine"
	"regexp"
)

//const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

const cityRe2 = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a></th></tr> <tr><td width="180"><span class="grayL">性别：</span>([^<]+)</td>`
const cityNextPageRe = `href="(http://www.zhenai.com/zhenghun/[^"]+)"`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe2)
	/*
		matchs := re.FindAll(contents, -1)
		matchs = matchs[6:]
		matchs = matchs[:len(matchs)-6]
		for _, m := range matchs{
			fmt.Printf("%s\n",m)
		}
	*/
	submatch := re.FindAllSubmatch(contents, -1)
	//for _, m := range submatch {
	//	fmt.Printf(">>>%s\n", m)
	//}
	result := engine.ParseResult{}
	for _, m := range submatch {
		name := string(m[2])
		gender := string(m[3])
		//result.Items = append(result.Items, "User "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, name, gender)
			},
		})
	}

	NextPageRe := regexp.MustCompile(cityNextPageRe)
	submatch1 := NextPageRe.FindAllSubmatch(contents, -1)
	for _, m := range submatch1 {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}
	return result
}
