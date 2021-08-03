package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	// application/x-www-form-urlencoded：为POST的contentType
	// strings.NewReader("mobile=xxxxxxxxxx&isRemberPwd=1") 理解为传递的参数
	resp, err := http.Post("http://localhost:8080/login.do",
		"application/x-www-form-urlencoded", strings.NewReader("mobile=xxxxxxxxxx&isRemberPwd=1"))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
