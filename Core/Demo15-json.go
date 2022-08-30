package main

import (
	"encoding/json"
	"fmt"
)

type Res struct {
	Name    string `json:"name"`
	Title   string `json:"title"`
	Contact struct {
		City  string `json:"city"`
		Phone string `json:"phone"`
	} `json:"contact"`
}

func main() {
	var JSON = `{
    "name": "Gopher",
    "title": "programmer",
    "contact": {
      "city": "beijing",
      "phone": "19999999999"
    }123
  }`
	fmt.Println("--------解析json字符串为对象--------------")
	var res Res
	json.Unmarshal([]byte(JSON), &res)
	fmt.Printf("结果：%+v\n", res)
	fmt.Printf("获取title= %v", res.Title)

	fmt.Println("--------解析json字符串为map集合--------------")
	var res2 map[string]interface{}
	err := json.Unmarshal([]byte(JSON), &res2)
	if err != nil {
		fmt.Println("json parser err:", err)
	}

	fmt.Printf("获取name=%v", res2["name"])

}
