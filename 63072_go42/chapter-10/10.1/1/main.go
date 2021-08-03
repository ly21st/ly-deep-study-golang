package main

import (
	"fmt"
)

type Human struct {
	name   string // 姓名
	Gender string // 性别
	Age    int    // 年龄
	string        // 匿名字段
}

func (h Human) print() { // 值方法
	fmt.Println("Human:", h)
}

type MyInt int

func (m MyInt) print() { // 值方法
	fmt.Println("MyInt:", m)
}

func main() {
	//使用new方式
	hu := new(Human)
	hu.name = "Titan"
	hu.Gender = "男"
	hu.Age = 14
	hu.string = "Student"
	hu.print()

	// 指针变量
	mi := new(MyInt)
	mi.print()

	// 使用结构体字面量赋值
	hum := Human{"Hawking", "男", 14, "Monitor"}
	hum.print()

	// 值变量
	myi := MyInt(99)
	myi.print()
}
