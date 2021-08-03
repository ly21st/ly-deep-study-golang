package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileObj, _ := os.OpenFile("./all.txt", os.O_RDWR|os.O_CREATE, 0666)
	defer fileObj.Close()

	Rd := bufio.NewReader(fileObj)
	cont, _ := Rd.ReadSlice('#')
	fmt.Println(string(cont))

	Wr := bufio.NewWriter(fileObj)
	Wr.WriteString("WriteString writes a ## string.")
	Wr.Flush()
}
