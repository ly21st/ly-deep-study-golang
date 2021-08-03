package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map

	// Store保存数据
	m.Store("name", "Joe")
	m.Store("gender", "Male")

	// LoadOrStore
	//若key不存在，则存入key和value，返回false和输入的value
	v, ok := m.LoadOrStore("name1", "Jim")
	fmt.Println(ok, v) //false Jim

	//若key已存在，则返回true和key对应的value，不会修改原来的value
	v, ok = m.LoadOrStore("name", "aaa")
	fmt.Println(ok, v) //true Joe

	// Load读取数据
	v, ok = m.Load("name")
	if ok {
		fmt.Println("key存在，值是： ", v)
	} else {
		fmt.Println("key不存在")
	}

	// Range 遍历
	// 遍历sync.Map
	f := func(k, v interface{}) bool {
		fmt.Println(k, v)
		return true
	}
	m.Range(f)

	// Delete 删除key及数据
	m.Delete("name1")
	fmt.Println(m.Load("name1"))

}
