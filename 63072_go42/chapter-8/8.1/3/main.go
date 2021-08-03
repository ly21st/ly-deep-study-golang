package main

import "fmt"

// Factorial函数递归调用
func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

// Fac2函数循环计算
func Fac2(n uint64) (result uint64) {
	result = 1
	var un uint64 = 1
	for i := un; i <= n; i++ {
		result *= i
	}
	return
}

func main() {
	var i uint64 = 7
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(i))
	fmt.Printf("%d 的阶乘是 %d\n", i, Fac2(i))

}
