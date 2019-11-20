package parser

import (
	"learngo/crawler/engine"
	"learngo/crawler/model"
	"regexp"
	"strconv"
	"strings"
)

//<div class="des f-cl" data-v-3c42fade>太原 | 35岁 | 大学本科 | 离异 | 165cm | 8001-12000元</div>
var personDataRe = regexp.MustCompile(`<div class="des f-cl" data-v-3c42fade>([^<]+)</div>`)
var personWeightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([\d]+)kg</div>`)
var userIdRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

func ParseProfile(contents []byte, url string, name string, gender string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	profile.Gender = gender
	match := personDataRe.FindSubmatch(contents)
	if match != nil {
		tmp := string(match[1])
		val := strings.Split(tmp, " | ")

		profile.Birthplace = val[0]

		val[1] = val[1][:len(val[1])-3]
		age, err := strconv.Atoi(string(val[1]))
		if err != nil {
			profile.Age = 0
		} else {
			profile.Age = age
		}

		profile.Education = val[2]
		profile.Marriage = val[3]

		val[4] = val[4][:len(val[4])-2]
		hi, err := strconv.Atoi(string(val[4]))
		if err != nil {
			profile.Height = 0
		} else {
			profile.Height = hi
		}
		profile.Income = val[5]
		//fmt.Printf("籍贯:%s， 年龄:%s, 教育:%s, 婚况:%s, 身高:%s, 收入:%s\n", val[0], val[1], val[2], val[3], val[4], val[5])
	}

	submatch := personWeightRe.FindSubmatch(contents)
	if submatch != nil {
		//fmt.Printf("weight=%skg\n", submatch[1])
		val, err := strconv.Atoi(string(submatch[1]))
		if err != nil {
			profile.Weight = 0
		} else {
			profile.Weight = val
		}
	} else {
		profile.Weight = 0
	}

	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Type:    "zhenai",
				Id:      extractString([]byte(url), userIdRe),
				Payload: profile,
			},
		},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	//fmt.Printf(">>>>>>%s", match)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
