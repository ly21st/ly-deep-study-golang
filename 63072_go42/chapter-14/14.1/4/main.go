package main

import (
	"fmt"
	"reflect"
)

func main() {
	t := 9
	// 反射创建int 变量
	varType := reflect.TypeOf(t)

	v1 := reflect.New(varType)
	v1.Elem().SetInt(1)
	varNew := v1.Elem().Interface()
	fmt.Printf("int Var: 指针：%d 值：%d\n", v1, varNew)

	// 反射创建map slice
	newSlice := make([]int, 5)
	newmap := make(map[string]int)
	sliceType := reflect.TypeOf(newSlice)
	mapType := reflect.TypeOf(newmap)

	// 创建新值
	ReflectSlice := reflect.MakeSlice(sliceType, 5, 5)
	Reflectmap := reflect.MakeMap(mapType)

	// 使用新创建的变量
	V := 99
	SliceV := reflect.ValueOf(V)
	ReflectSlice = reflect.Append(ReflectSlice, SliceV)
	intSlice := ReflectSlice.Interface().([]int)
	fmt.Println("Slice:", intSlice)

	Key := "Rose"
	Value := 999
	MapKey := reflect.ValueOf(Key)
	MapValue := reflect.ValueOf(Value)
	Reflectmap.SetMapIndex(MapKey, MapValue)
	mapStringInt := Reflectmap.Interface().(map[string]int)
	fmt.Println("Map:", mapStringInt)
}
