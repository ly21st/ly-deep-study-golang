package main

import (
	"fmt"
)

func div(a, b int) {

	defer func() {

		if r := recover(); r != nil {
			fmt.Printf("捕获到错误：%s\n", r)
		}
	}()

	if b < 0 {

		panic("除数需要大于0")
	}

	fmt.Println("余数为：", a/b)

}

func main() {
	// 捕捉内部的Panic错误
	div(10, 0)

	// 捕捉主动Panic的错误
	div(10, -1)
}
