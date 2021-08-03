package main

import "fmt"

const s = "Go语言"

func main() {

	for i, r := range s {
		fmt.Printf("%#U  ： %d\n", r, i)
	}
}
