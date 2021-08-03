package main

import (
	"fmt"
)

func main() {
	defer func() {
		fmt.Println("done")
		// 即使有panic，Println也正常执行。
		if err := recover(); err != nil {
			fmt.Printf("run time panic: %v \n", err)
		}
	}()
	fmt.Println("start")
	panic("Error") //   发生运行时异常的地方
}
