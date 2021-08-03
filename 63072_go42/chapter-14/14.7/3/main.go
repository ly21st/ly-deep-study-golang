package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fileObj, err := os.Open("./all.txt")
	defer fileObj.Close()

	contents, _ := ioutil.ReadAll(fileObj)
	fmt.Println(string(contents))

	if contents, _ := ioutil.ReadFile("./all.txt"); err == nil {
		fmt.Println(string(contents))
	}

	ioutil.WriteFile("./t3.txt", contents, 0666)

}
