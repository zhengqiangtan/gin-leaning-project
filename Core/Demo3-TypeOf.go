package main

import (
	"fmt"
	"reflect"
)

/**
类型断言测试
*/
func main() {

	var i interface{}

	var a float64 = 6.32
	i = a //空接口可以接收任意数据类型
	// 使用类型断言，检测机制，避免panic错误
	b, flag := i.(float64)
	if flag {
		fmt.Printf("b的类型为%T 值为%v", b, b) // b的类型为float64 值为6.32
	} else {
		fmt.Println("类型转换失败")
	}

	// 判断类型1
	JudgeType(i, 0, "string", 0.23, true)

	// 类型断言2
	var x = "张三"
	v2 := reflect.TypeOf(x) //类型
	fmt.Printf("%v", v2)    //输出string

}

func JudgeType(items ...interface{}) {
	for index, value := range items {
		switch value.(type) {
		case bool:
			fmt.Printf("第%v个参数时bool类型，值为%v\n", index, value)
		case float32:
			fmt.Printf("第%v个参数时float32类型，值为%v\n", index, value)
		case float64:
			fmt.Printf("第%v个参数时float64类型，值为%v\n", index, value)
		case int:
			fmt.Printf("第%v个参数时int类型，值为%v\n", index, value)
		case string:
			fmt.Printf("第%v个参数时string类型，值为%v\n", index, value)
		default:
			fmt.Printf("第%v个参数类型不确定，值为%v\n", index, value)
		}

	}
}
