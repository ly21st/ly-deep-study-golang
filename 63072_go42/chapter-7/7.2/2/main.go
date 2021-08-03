package main

import "fmt"

func main() {
	// 规则二 defer执行顺序为先进后出
	defer fmt.Print(" !!! ")
	defer fmt.Print(" world ")
	fmt.Print(" hello ")

}
