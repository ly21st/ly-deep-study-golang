package main

import (
	"fmt"
)

type I interface {
	f()
}

type T string

func (t T) f() {
	fmt.Println("T Method")
}

type Stringer interface {
	String() string
}

func main() {

	// 类型断言
	var varI I
	varI = T("Tstring")
	if v, ok := varI.(T); ok { // 类型断言
		fmt.Println("varI类型断言结果为：", v) // varI已经转为T类型
		varI.f()
	}

	// Type-switch做类型判断
	var value interface{} // 默认为零值

	switch str := value.(type) {
	case string:
		fmt.Println("value类型断言结果为string:", str)

	case Stringer:
		fmt.Println("value类型断言结果为Stringer:", str)

	default:
		fmt.Println("value类型不在上述类型之中")
	}

	// Comma-ok断言
	value = "类型断言检查"
	str, ok := value.(string)
	if ok {
		fmt.Printf("value类型断言结果为：%T\n", str) // str已经转为string类型
	} else {
		fmt.Printf("value不是string类型 \n")
	}
}
