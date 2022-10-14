package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//mytest()

	//case1_struct_to_json()

	//case2_map_to_json()
	//
	//case3_json_to_struct()
	//
	case4_json_to_map()

}

type Res struct {
	Name    string `json:"name"`
	Title   string `json:"title"`
	Contact struct {
		City  string `json:"city"`
		Phone string `json:"phone"`
	} `json:"contact"`
}

func mytest() {
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

//成员变量名首字母必须大写
type IT struct {
	Company  string   `json:"-"`        //此字段不会输出到屏幕
	Subjects []string `json:"subjects"` //二次编码
	IsOk     bool     `json:",string"`
	Price    float64  `json:",string"`
}

func case1_struct_to_json() {
	//定义一个结构体变量，同时初始化
	s := IT{"itcast", []string{"Go", "C++", "Python", "Test"}, true, 666.666}
	//编码，根据内容生成json文本
	//{"Company":"itcast","Subjects":["Go","C++","Python","Test"],"IsOk":true,"Price":666.666}
	//buf, err := json.Marshal(s)
	buf, err := json.MarshalIndent(s, "", " ") //格式化编码
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println(string(buf))
}

func case2_map_to_json() {
	//创建一个map
	m := make(map[string]interface{}, 4)
	m["company"] = "itcast"
	m["subjects"] = []string{"Go", "C++", "Python", "Test"}
	m["isok"] = true
	m["price"] = 666.666
	//编码成json
	//result, err := json.Marshal(m)
	result, err := json.MarshalIndent(m, "", "")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("result = ", string(result))
}

type IT2 struct {
	Company  string   `json:"company"`
	Subjects []string `json:"subjects"` //二次编码
	IsOk     bool     `json:"isok"`
	Price    float64  `json:"price"`
}

func case3_json_to_struct() {
	jsonBuf := `
    {
    "company": "itcast",
    "subjects": [
        "Go",
        "C++",
        "Python",
        "Test"
    ],
    "isok": true,
    "price": 666.666
	}`
	var tmp IT2                                  //定义一个结构体变量
	err := json.Unmarshal([]byte(jsonBuf), &tmp) //第二个参数要地址传递
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	//fmt.Println("tmp = ", tmp)
	fmt.Printf("tmp = %+v\n", tmp)

	type IT2 struct {
		Subjects []string `json:"subjects"` //二次编码
	}
	var tmp2 IT2
	err = json.Unmarshal([]byte(jsonBuf), &tmp2) //第二个参数要地址传递
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Printf("tmp2 = %+v\n", tmp2)
}

//tmp = {Company:itcast Subjects:[Go C++ Python Test] IsOk:true Price:666.666}
//tmp2 = {Subjects:[Go C++ Python Test]}

func case4_json_to_map() {
	jsonBuf := `
    {
    "company": "itcast",
    "subjects": [
        "Go",
        "C++",
        "Python",
        "Test"
    ],
    "isok": true,
    "price": 666.666
	}`

	//创建一个map
	m := make(map[string]interface{}, 4)
	err := json.Unmarshal([]byte(jsonBuf), &m) //第二个参数要地址传递
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Printf("m = %+v\n", m)
	//  var str string
	//  str = string(m["company"]) //err， 无法转换
	//  fmt.Println("str = ", str)

	var str string
	//类型断言, 值，它是value类型
	for key, value := range m {
		//fmt.Printf("%v ============> %v\n", key, value)
		switch data := value.(type) {
		case string:
			str = data
			fmt.Printf("map[%s]的值类型为string, value = %s\n", key, str)
		case bool:
			fmt.Printf("map[%s]的值类型为bool, value = %v\n", key, data)
		case float64:
			fmt.Printf("map[%s]的值类型为float64, value = %f\n", key, data)
		case []string:
			fmt.Printf("map[%s]的值类型为[]string, value = %v\n", key, data)
		case []interface{}:
			fmt.Printf("map[%s]的值类型为[]interface, value = %v\n", key, data)
		}
	}
}

//m = map[company:itcast isok:true price:666.666 subjects:[Go C++ Python Test]]
//map[company]的值类型为string, value = itcast
//map[subjects]的值类型为[]interface, value = [Go C++ Python Test]
//map[isok]的值类型为bool, value = true
//map[price]的值类型为float64, value = 666.666000
