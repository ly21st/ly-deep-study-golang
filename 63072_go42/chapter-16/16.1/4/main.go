package main

import (
	"encoding/json"
	"fmt"
)

type Human struct {
	Name   json.RawMessage `json:"name"` // 姓名，json.RawMessage 类型不会进行解码
	Gender string          `json:"s"`    // 性别，性别的tag表明在JSON中为s字段
	Age    int             `json:"Age"`  // 年龄
	Lesson
}

type Lesson struct {
	Lessons []string `json:"lessons"`
}

func main() {
	jsonStr := `{"Age": 18,"name": "Jim" ,"s": "男",
	"lessons":["English","History"],"Room":201,"n":null,"b":false}`

	var hu Human
	if err := json.Unmarshal([]byte(jsonStr), &hu); err == nil {
		fmt.Printf("\n 结构体Human \n")
		fmt.Printf("%+v \n", hu) // 可以看到Name字段未解码，还是字节数组
	}

	// 对延迟解码的Human.Name进行反序列化
	var UName string
	if err := json.Unmarshal(hu.Name, &UName); err == nil {
		fmt.Printf("\n Human.Name: %s \n", UName)
	}
}
