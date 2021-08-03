package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
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
	// JSON数据的字符串
	jsonStr := `{"Age": 18,"name": "Jim" ,"s": "男",
	"lessons":["English","History"],"Room":201,"n":null,"b":false}`
	strR := strings.NewReader(jsonStr)
	h := &Human{}

	// Decode 解码JSON数据到结构体Human中
	err := json.NewDecoder(strR).Decode(h)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(h)

	// 定义Encode需要的Writer
	f, err := os.Create("./t.json")

	// 把保存数据的Human结构体对象编码为JSON保存到文件
	json.NewEncoder(f).Encode(h)

}
