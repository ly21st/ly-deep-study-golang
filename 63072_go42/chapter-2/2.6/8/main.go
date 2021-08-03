package main

import (
	"fmt"
)

var x int = 10

func main() {
	goto TL
	fmt.Println(x)
TL:
	fmt.Println("TL")
}
