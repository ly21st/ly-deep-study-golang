package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	for k, v := range os.Args {

		fmt.Println(k, ":", v)
	}
}
