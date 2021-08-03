package main

import (
	"fmt"
)

type A struct {
	Face int
}
type Aa = A // 类型别名

func (a A) f() {
	fmt.Println("hi ", a.Face)
}

func main() {
	var s A = A{Face: 9}
	s.f()

	var sa Aa = Aa{Face: 9}
	sa.f()
}
