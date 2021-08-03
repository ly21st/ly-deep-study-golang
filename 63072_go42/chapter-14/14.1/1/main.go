package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	name string
}

type MyInt int

func main() {

	var a int = 9
	v := reflect.ValueOf(a)                         // 返回Value类型对象，值为9
	t := reflect.TypeOf(a)                          // 返回Type类型对象，值为int
	fmt.Println(v, t, v.Type(), v.Kind(), t.Kind()) // kind()返回底层基础类型

	var mi MyInt = 99
	mv := reflect.ValueOf(mi)                            // 返回Value类型对象，值为99
	mt := reflect.TypeOf(mi)                             // 返回Type类型对象，值为MyInt
	fmt.Println(mv, mt, mv.Type(), mv.Kind(), mt.Kind()) // kind()返回底层基础类型

	var b [5]int = [5]int{5, 6, 7, 8}
	fmt.Println(reflect.TypeOf(b), reflect.TypeOf(b).Kind(), reflect.TypeOf(b).Elem()) // [5]int array int

	var Pupil Student
	p := reflect.ValueOf(Pupil) // 使用ValueOf()获得结构体的Value对象

	fmt.Println(p.Type()) // 输出:Student
	fmt.Println(p.Kind()) // 输出:struct

}
