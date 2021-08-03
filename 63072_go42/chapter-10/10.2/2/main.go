package main

import (
	"fmt"
)

type T struct {
	Name string
}

func (t T) M1() {
	t.Name = "name1"
}

func (t *T) M2() {
	t.Name = "name2"
}
func main() {

	t2 := &T{"t2"}

	fmt.Println("M1调用前：", t2.Name)
	t2.M1()
	fmt.Println("M1调用后：", t2.Name)

	fmt.Println("M2调用前：", t2.Name)
	t2.M2()
	fmt.Println("M2调用后：", t2.Name)

}
