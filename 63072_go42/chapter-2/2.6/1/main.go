package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("false")
		fallthrough
	case true:
		fmt.Println("true")
		fallthrough
	case false:
		fmt.Println("false fallthrough")
		fallthrough
	case true:
		fmt.Println("true fallthrough")
		fallthrough
	default:
		fmt.Println("default")
	}
}
