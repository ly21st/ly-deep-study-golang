package main

import (
	"fmt"
)

func main() {
	callback(1, Add)
}

func Add(a, b int) {
	fmt.Printf("%d 与 %d 相加的和是: %d\n", a, b, a+b)
}

func callback(y int, f func(int, int)) {
	f(y, 2) // 回调函数f
}
