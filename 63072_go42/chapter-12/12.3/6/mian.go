package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once

func main() {

	for i, v := range make([]string, 10) {
		once.Do(onces)
		fmt.Println("v:", v, "---i:", i)
	}

	for i := 0; i < 10; i++ {

		go func(i int) {
			once.Do(onced)
			fmt.Println(i)
		}(i)
	}
	time.Sleep(10)
}
func onces() {
	fmt.Println("once")
}
func onced() {
	fmt.Println("onced")
}
