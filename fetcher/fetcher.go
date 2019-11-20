package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
)

const zhenai_url = "http://www.zhenai.com/zhenghun"
const province_url = "http://www.maps7.com/china_province.php"
const huazhengcaiwu = "https://www.huazhengcaiwu.com/city/"

var rateLimiter = time.Tick(450 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	<-rateLimiter
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_4) AppleWebKit/537.31 (KHTML, like Gecko) Chrome/71.0.3578.87 Safari/537.21")

	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			//fmt.Println("Redirect:", req)
			return nil
		},
	}
	resp, err := client.Do(req)

	//resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Wrong status code: %d", resp.StatusCode)
	}
	/*
		//如果读取的网页不是utf-8编码，中文就会显示乱码，此部分注释的代码就是转码操作
		bodyReader := bufio.NewReader(resp.Body)
		e := determineEncoding(bodyReader)
		utf8Read := transform.NewReader(bodyReader,e.NewDecoder())
		return ioutil.ReadAll(utf8Read)
	*/

	return ioutil.ReadAll(resp.Body)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
