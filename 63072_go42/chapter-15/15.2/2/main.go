package main

import (
	"log"
	"os"
	"text/template"
)

var report = template.Must(template.ParseFiles("tmp.txt"))

type Book struct {
	Name  string
	Price float64
}

func main() {
	Data := []Book{{"《三国演义》", 19.82}, {"《儒林外史》", 99.09}, {"《史记》", 26.89}}
	if err := report.Execute(os.Stdout, Data); err != nil {
		log.Fatal(err)
	}
}
