package main

import (
	"fmt"
)

type MyInt int
type NewInt = MyInt

func (m MyInt) print() { // 值方法
	fmt.Println("MyInt:", m)
}

func main() {
	myi := MyInt(99)
	myi.print()

	Ni := NewInt(myi)
	Ni.print()
}
