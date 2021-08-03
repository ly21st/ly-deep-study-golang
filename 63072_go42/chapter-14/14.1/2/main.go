package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	name string
	Age  int
}

func main() {
	var a int = 99
	v := reflect.ValueOf(a) // 返回Value类型对象，值为9
	t := reflect.TypeOf(a)  // 返回Type类型对象，值为int

	fmt.Println(v.Type(), t.Kind(), reflect.ValueOf(&a).Elem())
	fmt.Println(reflect.ValueOf(a).CanSet(), reflect.ValueOf(a).CanAddr())
	fmt.Println(reflect.ValueOf(&a).CanSet(), reflect.ValueOf(&a).CanAddr())

	pa := reflect.ValueOf(&a).Elem() // reflect.Value.Elem()表示获取原始值对应的反射对象
	fmt.Println("CanSet:", pa.CanSet(), "CanAddr:", pa.CanAddr())
	//fmt.Println(reflect.ValueOf(a).Elem().CanSet(), reflect.ValueOf(a).Elem().CanAddr())
	fmt.Println(pa, pa.CanSet())

	pa.SetInt(100)
	fmt.Println(pa)

	var Pupil Student = Student{"Jim", 8}
	Pupilv := reflect.ValueOf(Pupil) // 使用ValueOf()获取结构体的Value对象

	fmt.Println(Pupilv.Type()) // 输出:Student
	fmt.Println(Pupilv.Kind()) // 输出:struct

	p := reflect.ValueOf(&Pupil).Elem() // 获取原始值对应的反射对象
	fmt.Println(reflect.ValueOf(&Pupil))
	fmt.Println("CanSet:", p.CanSet(), "CanAddr:", p.CanAddr())

	//p.Field(0).SetString("Mike") // 未导出字段，不能修改，会发生异常
	p.Field(1).SetInt(10)
	fmt.Println(p)

}
