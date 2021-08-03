package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// var DefaultClient = &Client{}
	// func Get(url string) (resp *Response, err error) {
	// return DefaultClient.Get(url)
	// }
	/*
		func (c *Client) Get(url string) (resp *Response, err error) {
			req, err := NewRequest("GET", url, nil)
			if err != nil {
				return nil, err
			}
			return c.Do(req)
		}
	*/

	// http.Get实际上是DefaultClient.Get(url)，Get函数是高度封装的，只有一个参数url。
	// 对于一般的http Request是可以使用，但是不能定制Request。
	response, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println(err)
	}

	// 程序在使用完回复后必须关闭回复的主体。
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}
