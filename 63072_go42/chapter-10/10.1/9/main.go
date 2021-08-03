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

type Student struct {
	Human     // 匿名字段
	Room  int // 教室
	int       // 匿名字段
}

func (h Human) String() { // 值方法
	fmt.Println("Human")
}

func (s Student) String() { // 值方法
	fmt.Println("Student")
}

func (s Student) Print() { // 值方法
	fmt.Println("Print")
}

func main() {
	stud := Student{Room: 102, Human: Human{"Hawking", "男", 14, "Monitor"}}
	stud.String()
	stud.Human.String()
}
