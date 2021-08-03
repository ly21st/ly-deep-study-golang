package main

import "fmt"

func main() {
	data := []int{1, 2, 3}
	for _, v := range data {
		v *= 10 // 通常数据项不会改变
	}
	fmt.Println("data:", data) // 程序输出: [1 2 3]
}
