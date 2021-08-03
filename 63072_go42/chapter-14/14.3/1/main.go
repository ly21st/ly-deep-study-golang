package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{33, 5, 3, -6, 19, 11, -14}
	sort.Ints(a)
	fmt.Println("排序：", a)

	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	fmt.Println("降序: ", a)

	s := []string{"surface", "ipad", "Lenovo", "mac", "thinkpad", "联想"}
	sort.Strings(s)
	fmt.Println("排序：", s)

	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Printf("降序: %v\n", s)
}
