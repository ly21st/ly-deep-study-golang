package main

import (
	"fmt"
)

type MyInt int

func (mi *MyInt) print() { // 指针接收器，指针方法
	fmt.Println("MyInt:", *mi)
}
func (mi MyInt) echo() { // 值接收器，值方法
	fmt.Println("MyInt:", mi)
}

func main() {
	i := MyInt(9)
	i.print()
}
