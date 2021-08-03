package main

import (
	"fmt"
)

func main() {
	fn := func() {
		fmt.Println("hello")
	}
	fn()

	fmt.Println("匿名函数加法求和：", func(x, y int) int { return x + y }(3, 4))

	func() {
		sum := 0
		for i := 1; i <= 1e6; i++ {
			sum += i
		}
		fmt.Println("匿名函数加法循环求和：", sum)
	}()
}
