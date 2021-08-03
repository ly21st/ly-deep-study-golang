package main

import (
	"fmt"
	"time"
)

func main() {
	data := []string{"one", "two", "three"}

	for _, v := range data {
		go func(in string) {
			fmt.Println(in)
		}(v)
	}

	time.Sleep(3 * time.Second)
	// goroutine输出: one, two, three
}
