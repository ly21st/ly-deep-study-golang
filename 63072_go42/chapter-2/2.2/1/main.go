package main

import (
	"fmt"
)

var x int = 10

func main() {
	x := 1         // 这里全局变量x 和 简式变量 x 会有变量隐藏，全局变量x在这层失效
	fmt.Println(x) // 显示 1，简式变量在这层block生效
	{
		fmt.Println(x) // 显示 1
		x := 2
		fmt.Println(x) // 显示 2
	}
	fmt.Println(x) // 注意：显示 1 (不是2)
}
