package main

import (
	"fmt"
	"time"
)

type funcType func(time.Time) // 定义函数类型funcType

func main() {
	f := func(t time.Time) time.Time { return t } // 直接赋值给变量
	fmt.Println(f(time.Now()))

	var timer funcType = CurrentTime // 定义函数类型funcType变量timer
	timer(time.Now())

	funcType(CurrentTime)(time.Now()) // 先把CurrentTime函数转为funcType类型，然后传入参数调用
	// 这种处理方式在Go 中比较常见
}

func CurrentTime(start time.Time) {
	fmt.Println(start)
}
