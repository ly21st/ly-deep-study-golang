package main

import "fmt"

func get() []byte {
	raw := make([]byte, 10000)
	fmt.Println(len(raw), cap(raw), &raw[0]) // prints: 10000 10000 <byte_addr_x>
	return raw[:3]                           // 10000个字节实际只需要引用3个，其他空间浪费
}

func main() {
	data := get()
	fmt.Println(len(data), cap(data), &data[0]) // prints: 3 10000 <byte_addr_x>
}
