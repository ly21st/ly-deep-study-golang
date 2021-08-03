package main

import "fmt"

func main() {
	var i int = 1

	// 规则一，当defer被声明时，其后面函数参数会被实时解析
	// 注意，fmt.Println在defer后面，它的参数会实时计算
	// 输出: result => 2 (而不是 4)
	defer fmt.Println("result1 =>", func() int { return i * 2 }())
	i++

	// 下面defer后面的函数无参数，所以最里层的i应该是3
	defer func() {
		fmt.Println("result2 =>", i*2)
	}()
	i++
}
