package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	//citylist_test_data.html文件内容是https://www.huazhengcaiwu.com/city/网址的源代码
	if err != nil {
		panic(err)
	}

	results := ParseCityList(contents)
	const resultSize = 301
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/beijing/",
		"http://www.zhenai.com/zhenghun/shanghai/",
		"http://www.zhenai.com/zhenghun/guangzhou/",
	}

	if len(results.Requests) != resultSize {
		t.Errorf("result should have %d requests; but had %d", resultSize, len(results.Requests))
	}

	for i, url := range expectedUrls {
		if results.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but was %s", i, url, results.Requests[i].Url)
		}
	}
}
