package main

import "fmt"

var G int = 7

func main() {
	// 影响全局变量G，代码块状态持续
	y := func() int {
		fmt.Printf("G: %d, G的地址:%p\n", G, &G)
		G += 1
		return G
	}
	fmt.Println(y(), y)
	fmt.Println(y(), y)
	fmt.Println(y(), y) //y的地址

	// 影响全局变量G，注意z的匿名函数是直接执行，所以结果不变
	z := func() int {
		G += 1
		return G
	}()
	fmt.Println(z, &z)
	fmt.Println(z, &z)
	fmt.Println(z, &z)

	// 影响外层（自由）变量i，代码块状态持续
	var f = N()
	fmt.Println(f(1), &f)
	fmt.Println(f(1), &f)
	fmt.Println(f(1), &f)

	var f1 = N()
	fmt.Println(f1(1), &f1)

}

func N() func(int) int {
	var i int
	return func(d int) int {
		fmt.Printf("i: %d, i的地址:%p\n", i, &i)
		i += d
		return i
	}
}
