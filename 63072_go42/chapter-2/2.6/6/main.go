package main

import (
	"fmt"
)

func main() {

	fmt.Println("start")
OuterLoop:
	for i := 0; i < 3; i++ {
		fmt.Println("i循环:", i)
		for j := 0; j < 3; j++ {
			fmt.Println("j循环:", j)
			switch j {
			case 0:
				fmt.Println("break")
				break
			case 2:
				fmt.Println("2 OuterLoop")
				break OuterLoop
			}
			fmt.Println("switch:", j)
		}
	}

	fmt.Println("end")
}
