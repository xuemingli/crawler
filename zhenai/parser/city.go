package parser

import (
	"learngo/crawler/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	/*
		matchs := re.FindAll(contents, -1)
		matchs = matchs[6:]
		matchs = matchs[:len(matchs)-6]
		for _, m := range matchs{
			fmt.Printf("%s\n",m)
		}
	*/
	submatch := re.FindAllSubmatch(contents, -1)
	//submatch = submatch[6:]
	//submatch = submatch[:len(submatch)-6]
	//var w string
	result := engine.ParseResult{}
	for _, m := range submatch {
		name := string(m[2])
		result.Items = append(result.Items, "User "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, name)
			}})
		//fmt.Printf("city: %s, url: %s\n", m[2], w)
	}
	return result
}
