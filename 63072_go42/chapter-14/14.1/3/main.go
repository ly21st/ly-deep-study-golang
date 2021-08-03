package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	name string
	Age  int `json:"years"`
}

func main() {
	var Pupil Student = Student{"joke", 18}
	setStudent := reflect.ValueOf(&Pupil).Elem()

	sSAge, _ := setStudent.Type().FieldByName("Age")
	fmt.Println(sSAge.Tag.Get("json")) // years
}
