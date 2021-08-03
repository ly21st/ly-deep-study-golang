package main

import (
	"fmt"
)

type printer interface {
	print()
}

func (p printer) print() { //  invalid receiver type printer (printer is an interface type)
	fmt.Println("printer:", p)
}
func main() {}
