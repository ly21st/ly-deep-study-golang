package main

func main() {
	var m map[string]int
	m["one"] = 1 // panic: assignment to entry in nil map

}
