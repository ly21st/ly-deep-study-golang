package main

import (
	"fmt"
)

type MyInt int

func (m MyInt) print() { // 值方法
	fmt.Println("MyInt:", m)
}

func main() {
	myi := MyInt(99)
	myi.print()
}
