package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var ss = make(chan int)

func main() {
	go signalListen()
	select {
	case <-ss:
		break
	}
}

func signalListen() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGQUIT)
	for {
		s := <-c
		//收到信号后的处理，这里只是输出信号内容，可以做一些更有意思的事
		fmt.Println("get signal:", s)
		ss <- 9
		break
	}
}
