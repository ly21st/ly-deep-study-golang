package main

import (
	"fmt"
	"unsafe"
)

type V struct {
	b byte
	i int32
	j int64
}

func (v V) GetI() {
	fmt.Printf("i=%d\n", v.i)
}
func (v V) GetJ() {
	fmt.Printf("j=%d\n", v.j)
}

func main() {
	// 定义指针类型变量
	var v *V = new(V)

	// v的长度
	fmt.Printf("size=%d\n", unsafe.Sizeof(*v))
	// 取得v的指针考虑对齐值计算偏移量，然后转为*int32的值，对应结构体的i。
	var i *int32 = (*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(v)) + uintptr(4*unsafe.Sizeof(byte(0)))))

	fmt.Println("指针地址：", i)
	fmt.Println("指针uintptr值:", uintptr(unsafe.Pointer(i)))
	*i = int32(98)

	// 根据v的基准地址加上偏移量进行指针运算，运算后的值为j的地址，使用unsafe.Pointer转为指针
	var j *int64 = (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(v)) + uintptr(unsafe.Sizeof(int64(0)))))

	*j = int64(763)
	fmt.Println("指针uintptr值:", uintptr(unsafe.Pointer(&v.b)))
	fmt.Println("指针uintptr值:", uintptr(unsafe.Pointer(&v.i)))
	fmt.Println("指针uintptr值:", uintptr(unsafe.Pointer(&v.j)))
	v.GetI()
	v.GetJ()
}
