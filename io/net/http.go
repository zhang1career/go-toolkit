package net

import (
	"bufio"
	"fmt"
	"github.com/zhang1career/app/www/action/file"
	"github.com/zhang1career/lib/log"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
)

var defaultHeaders = map[string]string{
	"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_4) AppleWebKit/537.31 (KHTML, like Gecko) Chrome/71.0.3578.87 Safari/537.21",
}

// param
//   url
//   opts...
//     [0]: http method
//     [1]: http headers
func Curl(url string, opts ...interface{}) ([]byte, error) {
	method := http.MethodGet
	if len(opts) > 0 {
		method = opts[0].(string)
	}
	headers := defaultHeaders
	if len(opts) > 1 {
		headers = opts[1].(map[string]string)
	}
		
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			log.Info("redirect:%s", req)
			return nil
		},
	}
	req, err := http.NewRequest(method, url, nil)
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Error(err.Error())
		}
	}()
	
	if resp.StatusCode != http.StatusOK {
		log.Warn("status:%d", resp.StatusCode)
	}
	encodingReader := bufio.NewReader(resp.Body)
	e := determineEncoding(encodingReader)
	utf8Reader := transform.NewReader(resp.Body, e.NewEncoder())
	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding (r *bufio.Reader) (ret encoding.Encoding) {
	ret = unicode.UTF8
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Error(err.Error())
		return ret
	}
	ret, _, _ = charset.DetermineEncoding(bytes, "")
	return ret
}

func Get(url string) ([]byte, error) {
	return Curl(url)
}

func Post(url string) ([]byte, error) {
	return Curl(url, http.MethodPost)
}


func Client() {
	ret, err := Curl("localhost:9000/")
	if err != nil {
		log.Error(err.Error())
	}
	
	fmt.Printf("%s\n", ret)
}

func Server() {
	act := file.Action{}
	http.HandleFunc("/", WrapError(&act))
	err := http.ListenAndServe(":" + Port, nil)
	if err != nil {
		panic(err)
	}
}