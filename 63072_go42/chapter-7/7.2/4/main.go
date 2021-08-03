package main

import (
	"fmt"
)

func main() {
	fmt.Println("=========================")
	fmt.Println("fun1 return:", fun1())

	fmt.Println("=========================")
	fmt.Println("fun2 return:", fun2())

	fmt.Println("=========================")
	fmt.Println("fun3 return:", fun3())

	fmt.Println("=========================")
	fmt.Println("fun4 return:", fun4())
}

func fun1() (i int) {
	defer func() {
		i++
		fmt.Println("fun1 defer2:", i) // 打印结果为 fun1 defer2: 2
	}()

	// 规则二 defer执行顺序为先进后出
	defer func() {
		i++
		fmt.Println("fun1 defer1:", i) // 打印结果为 fun1 defer1: 1
	}()

	// 规则三 defer可以读取有名返回值（函数指定了返回参数名）
	return 0 //这里实际结果为2。如果是return 100呢
}

func fun2() int {
	var i int
	defer func() {
		i++
		fmt.Println("fun2 defer2:", i) // 打印结果为 fun2 defer2: 2
	}()

	defer func() {
		i++
		fmt.Println("fun2 defer1:", i) // 打印结果为 fun2 defer1: 1
	}()
	return i
}

func fun3() (r int) {
	t := 5
	defer func() {
		t = t + 5
		fmt.Println("fun3 defer:", t) // 打印结果为 fun3 defer: 10
	}()
	return t
}

func fun4() int {
	i := 8
	// 规则一 当defer被声明时，其参数会被实时解析
	defer func(i int) {
		fmt.Println("fun4 defer:", i) // 打印结果为 fun4 defer: 8
	}(i)
	i = 19
	return i
}
