package main

import "fmt"

func main() {
	var a, b int = 20, 30 // 声明实际变量
	var ptra *int         // 声明指针变量
	var ptrb *int = &b

	ptra = &a // 指针变量的存储地址

	fmt.Printf("a  变量的地址是: %x\n", &a)
	fmt.Printf("b  变量的地址是: %x\n", &b)

	// 指针变量的存储地址
	fmt.Printf("ptra  变量的存储地址: %x\n", ptra)
	fmt.Printf("ptrb  变量的存储地址: %x\n", ptrb)

	// 使用指针访问值
	fmt.Printf("*ptra  变量的值: %d\n", *ptra)
	fmt.Printf("*ptrb  变量的值: %d\n", *ptrb)
}
