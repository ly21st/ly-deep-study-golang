package main

import (
	"fmt"
)

func (i int) print() { // cannot define new methods on non-local type int
	fmt.Println("Int:", i)
}

func main() {
}
