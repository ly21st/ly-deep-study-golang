package main

import (
	"fmt"
)

type T struct {
	a int
}

func (tv T) Mv(a int) int {
	fmt.Printf("Mv的值是: %d\n", a)
	return a
} // 值方法

func (tp *T) Mp(f float32) float32 {
	fmt.Printf("Mp: %f\n", f)
	return f
} // 指针方法

func main() {
	var t T
	// 下面几种调用方法是等价的
	t.Mv(1)    // 一般调用
	T.Mv(t, 1) // 显式接收器t可以当做为函数的第一个参数
	f0 := t.Mv // 通过选择器（selector）t.Mv将方法值赋值给一个变量 f0
	f0(2)
	T.Mv(t, 3)
	(T).Mv(t, 4)
	f1 := T.Mv // 利用方法表达式(Method Expression) T.Mv 取到函数值
	f1(t, 5)
	f2 := (T).Mv // 利用方法表达式(Method Expression) T.Mv 取到函数值
	f2(t, 6)
}
