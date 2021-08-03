package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

func LockA() {
	mutex.Lock()
	fmt.Println("Lock in A")
	LockB()
	time.Sleep(5)
	fmt.Println("Wake up in A")
	mutex.Unlock()
	fmt.Println("UnLock in A")
}
func LockB() {
	fmt.Println("B")
	mutex.Lock()
	fmt.Println("Lock in B")
	mutex.Unlock()
	fmt.Println("UnLock in B")
}
func main() {
	LockA()
	time.Sleep(10)
}
