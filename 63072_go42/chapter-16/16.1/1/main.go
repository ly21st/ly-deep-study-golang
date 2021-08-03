package main

import (
	"encoding/json"
	"fmt"
)

type Human struct {
	name   string `json:"name"` // 姓名
	Gender string `json:"s"`    // 性别，性别的tag表明在JSON中为s字段
	Age    int    `json:"Age"`  // 年龄
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
		fmt.Println("\n结构体Human")
		fmt.Println(hu)
	}

	var le Lesson
	if err := json.Unmarshal([]byte(jsonStr), &le); err == nil {
		fmt.Println("\n结构体Lesson")
		fmt.Println(le)
	}

	jsonStr = `["English","History"]`

	var str []string
	if err := json.Unmarshal([]byte(jsonStr), &str); err == nil {
		fmt.Println("\n字符串数组")
		fmt.Println(str)
	} else {
		fmt.Println(err)
	}
}
