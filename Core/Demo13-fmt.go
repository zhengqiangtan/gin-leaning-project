package main

import (
	"fmt"
	"strconv"
)

type student struct {
	name string
	id   int32
}

/**
%v 只输出所有的值；
%+v 先输出字段名字，再输出该字段的值；
%#v 先输出结构体名字值，再输出结构体（字段名字+字段的值）；
%s — 字符串
%d — 10进制数值
%T — type(值)
*/
func main() {
	//a := &student{
	//	name: "zhangsan",
	//	id:   1,
	//}
	//fmt.Printf("a=%v \n", a)
	//fmt.Printf("a=%+v \n", a)
	//fmt.Printf("a=%#v \n", a)
	//
	//fmt.Println(unsafe.Sizeof(struct{}{})) //0

	const name, age = "Kim", 22
	fmt.Print(name, age) //Kim22 不带换行原样输出
	fmt.Println()
	fmt.Println(name, age) // 带换行+空格
	fmt.Printf("name=%s and age = %+v", name, age)

}

//a=&{zhangsan 1}
//a=&{name:zhangsan id:1}
//a=&main.student{name:"zhangsan", id:1}

func case1_strconv_1() {
	strInt := "100"
	parseBool, err2 := strconv.ParseBool("1")
	fmt.Println(parseBool, err2)

	num, err := strconv.Atoi(strInt)
	if err != nil {
		fmt.Println("can't convert to int", err)
		return
	}
	fmt.Printf("type:%T\nvalue:%#v\n", num, num)
	// type:int
	// value:100
}
