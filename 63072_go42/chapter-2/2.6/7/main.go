package main

import (
	"fmt"
)

var x int = 10

func main() {

	if x%2 == 1 {
		goto L1
	}
	for x < 10 {
		x--
		fmt.Println(x)
	L1:
		x--
		fmt.Println(x)
	}

}
