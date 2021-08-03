package main

import "fmt"

func fun1() (i int) {
	// 规则三，defer可以读取有名返回值（函数指定了返回参数名）
	defer func() {
		i = i + 10 // defer可以读取有名返回值
	}()

	return 0 // 一般会认为返回0，实际上是10
}

func main() {
	fmt.Println("result2 =>", fun1())
}
