package main

import "fmt"

func main() {
	x := [3]int{1, 2, 3}

	func(arr *[3]int) {
		(*arr)[0] = 7
		fmt.Println(arr) // 显示 &[7 2 3]
	}(&x)

	fmt.Println(x) // 显示 [7 2 3]
}
