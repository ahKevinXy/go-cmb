package help

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"strings"
)

const (
	POST  = "POST"
	GET   = "GET"
	PUT   = "PUT"
	PATCH = "PATCH"
)

// MakeHttpRequest
// @Description: HTTP请求
// @Author ahKevinXy
// @Date 2022-11-08 16:49:40
func MakeHttpRequest(method, url string, entity map[string]interface{}, jar *cookiejar.Jar) (string, int, error) {
	var body io.Reader
	var err error

	if entity != nil {
		switch method {
		case POST, PUT, PATCH:
			b, err := json.Marshal(entity)
			if err != nil {
				return "", 0, err
			}

			b = bytes.Replace(b, []byte("\\u003c"), []byte("<"), -1)
			b = bytes.Replace(b, []byte("\\u003e"), []byte(">"), -1)
			b = bytes.Replace(b, []byte("\\u0026"), []byte("&"), -1)

			body = bytes.NewBuffer(b)

		case GET:
			if len(entity) > 0 {
				params := make([]string, len(entity))
				index := 0
				for k, v := range entity {
					_v := fmt.Sprintf("%v", v)
					params[index] = fmt.Sprintf("%s=%v", k, _v)
					index++
				}
				queryStr := strings.Join(params, "&")
				url = fmt.Sprintf("%s?%s", url, queryStr)
			}
		}
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return "", 0, err
	}
	if entity != nil && (method == POST || method == PUT || method == PATCH) {
		req.Header.Set("Content-Type", "application/json;charset=utf-8")
		req.Header.Set("Accept", "application/json")
	}
	req.Header.Add("Connection", "close")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux i686; U;) Gecko/20070322 Kazehakase/0.4.5")

	//fmt.Printf("================Request: %+v\n", *req)
	client := http.DefaultClient
	if jar != nil {
		client = &http.Client{
			Jar: jar,
		}
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("faild to do the request with error ", err)
		return "", 0, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated && res.StatusCode != http.StatusNoContent {
		fmt.Println("code is not 200 ", res.StatusCode, res)
		return "", 0, errors.New("http request failed to call")
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("could not read the response body")
		return "", 0, errors.New("the response could not be read")
	}

	return string(resBody), res.StatusCode, nil
}
