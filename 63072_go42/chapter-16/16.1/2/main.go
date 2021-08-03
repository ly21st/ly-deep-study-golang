package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonStr := `{"Age": 18,"name": "Jim" ,"s": "男","Lessons":["English","History"],"Room":201,"n":null,"b":false}`

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &data); err == nil {
		fmt.Println("map结构")
		fmt.Println(data)
	}

	for k, v := range data {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "是string", vv)
		case bool:
			fmt.Println(k, "是bool", vv)
		case float64:
			fmt.Println(k, "是float64", vv)
		case nil:
			fmt.Println(k, "是nil", vv)
		case []interface{}:
			fmt.Println(k, "是array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "未知数据类型")
		}
	}
}
