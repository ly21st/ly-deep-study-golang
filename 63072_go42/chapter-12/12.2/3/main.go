package main

import (
	"fmt"
)

func main() {
	c1 := make(chan int, 1)
	c1 <- 100
	close(c1)
	fmt.Println(<-c1) // 通道关闭后可正常接收数据

	c1 <- 100 // 通道关闭后发送数据会panic
}
