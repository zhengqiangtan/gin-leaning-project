package main

import (
	"fmt"
	"unsafe"
)

type student struct {
	name string
	id   int32
}

/**
%v 只输出所有的值；
%+v 先输出字段名字，再输出该字段的值；
%#v 先输出结构体名字值，再输出结构体（字段名字+字段的值）；
*/
func main() {
	a := &student{
		name: "zhangsan",
		id:   1,
	}
	fmt.Printf("a=%v \n", a)
	fmt.Printf("a=%+v \n", a)
	fmt.Printf("a=%#v \n", a)

	fmt.Println(unsafe.Sizeof(struct{}{})) //0

}

//a=&{zhangsan 1}
//a=&{name:zhangsan id:1}
//a=&main.student{name:"zhangsan", id:1}
