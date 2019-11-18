package parser

import (
	"learngo/crawler/engine"
	"learngo/crawler/model"
	"regexp"
	"strings"
)

//<div class="des f-cl" data-v-3c42fade>太原 | 35岁 | 大学本科 | 离异 | 165cm | 8001-12000元</div>
var personDataRe = regexp.MustCompile(`<div class="des f-cl" data-v-3c42fade>([^<]+)</div>`)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name

	match := personDataRe.FindSubmatch(contents)

	//if match != nil {
	//	age, err := strconv.Atoi(string(match[1]))
	//	if err != nil {
	//		//get user age in age
	//		profile.Age = age
	//	}
	//}
	if match != nil {
		tmp := string(match[1])
		val := strings.Split(tmp, " | ")
		//fmt.Printf("籍贯:%s， 年龄:%s, 教育:%s, 婚况:%s, 身高:%s, 收入:%s\n", val[0], val[1], val[2], val[3], val[4], val[5])
		profile.Birthplace = val[0]
		profile.Age = val[1]
		profile.Education = val[2]
		profile.Marriage = val[3]
		profile.Height = val[4]
		profile.Income = val[5]
	}
	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}
