package main

type T struct {
	Name string
}
type Intf interface {
	M1()
	M2()
}

func (t T) M1() {
	t.Name = "name1"
}

func (t *T) M2() {
	t.Name = "name2"
}
func main() {

	var t1 T = T{"t1"}
	t1.M1()
	t1.M2()

	var t2 Intf = &t1
	t2.M1()
	t2.M2()
}
