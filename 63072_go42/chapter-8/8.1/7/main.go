package main

import "fmt"

func Greeting(who ...string) {
	for k, v := range who {

		fmt.Println(k, v)
	}
}

func main() {
	s := []string{"James", "Jasmine"}
	Greeting(s...) // 注意这里切片s... ，把切片打散传入，与s具有相同底层数组的值。
}
