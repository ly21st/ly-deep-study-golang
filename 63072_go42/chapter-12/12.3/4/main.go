package main

import (
	"fmt"
	"sync"
	"time"
)

var m *sync.RWMutex

func main() {
	wg := sync.WaitGroup{}
	wg.Add(20)
	var rwMutex sync.RWMutex
	Data := 0
	for i := 0; i < 10; i++ {
		go func(t int) {
			rwMutex.RLock()
			defer rwMutex.RUnlock()
			fmt.Printf("读数据: %v %d\n", Data, i)
			wg.Done()
			//time.Sleep(2 * time.Second)
			// 这句代码第一次运行后，读解锁。
			// 循环到第二个时，读锁定后，这个goroutine就没有阻塞，同时读成功。
		}(i)

		go func(t int) {
			rwMutex.Lock()
			defer rwMutex.Unlock()
			Data += t
			fmt.Printf("写数据: %v %d \n", Data, t)
			wg.Done()

			// 对读写锁进行读锁定或者写锁定，都将阻塞（sleep设置长点可以更好看到效果）。写锁定下是需要解锁后才能写的。
			time.Sleep(5 * time.Second)
		}(i)
	}
	wg.Wait()
}
