package main

import (
	"fmt"
)

type MyInt int

type Q *MyInt

func (q Q) print() { // invalid receiver type Q (Q is a pointer type)
	fmt.Println("Q:", q)
}

func main() {}
