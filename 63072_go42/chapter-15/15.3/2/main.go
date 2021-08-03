package main

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {
	// 简式声明一个http.Client空结构体指针对象
	client := &http.Client{}

	// 使用http.NewRequest构建http Request请求
	request, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	if err != nil {
		fmt.Println(err)
	}

	// 使用http.Cookie结构体初始化一个cookie键值对
	cookie := &http.Cookie{Name: "userId", Value: strconv.Itoa(12345)}

	// 使用前面构建的request方法AddCookie往请求中添加cookie
	request.AddCookie(cookie)

	// 设置request的Header，具体可参考http协议
	request.Header.Set("Accept", "text/html, application/xhtml+xml, application/xml;q=0.9, */*;q=0.8")
	request.Header.Set("Accept-Charset", "GBK, utf-8;q=0.7, *;q=0.3")
	request.Header.Set("Accept-Encoding", "gzip, deflate, sdch")
	request.Header.Set("Accept-Language", "zh-CN, zh;q=0.8")
	request.Header.Set("Cache-Control", "max-age=0")
	request.Header.Set("Connection", "keep-alive")

	// 使用http.Client 来发送request，这里使用了Do方法。
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 程序结束时关闭response.Body响应流
	defer response.Body.Close()

	// 接收到的http Response 状态值
	fmt.Println(response.StatusCode)
	if response.StatusCode == 200 { // 200意味成功得到http Server返回的http Response信息

		// gzip.NewReader对压缩的返回信息解压（考虑网络传输量，http Server一般都会对响应压缩后再返回）
		body, err := gzip.NewReader(response.Body)
		if err != nil {
			fmt.Println(err)
		}

		defer body.Close()

		r, err := ioutil.ReadAll(body)
		if err != nil {
			fmt.Println(err)
		}
		// 打印出http Server返回的http Response信息
		fmt.Println(string(r))
	}
}
