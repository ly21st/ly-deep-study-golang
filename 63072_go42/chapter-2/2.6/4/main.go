package main

import (
	"fmt"
	"time"
)

type field struct {
	name string
}

func (p *field) print() {
	fmt.Println(p.name)
}

func main() {
	data := []field{{"one"}, {"two"}, {"three"}}

	for _, v := range data {
		go v.print()
	}
	time.Sleep(3 * time.Second)
	// goroutines （可能）显示: three, three, three
}
