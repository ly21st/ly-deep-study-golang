package main

import (
	"fmt"
	"os"
)

func main() {
	b := make([]byte, 1024)
	f, err := os.Open("./tt.txt") // 只读模式打开文件
	_, err = f.Read(b)
	f.Close()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))

}
