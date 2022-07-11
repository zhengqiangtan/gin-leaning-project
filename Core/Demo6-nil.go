package main

import "fmt"

/**
nil 只能赋值给指针、chan、func、interface、map 或 slice 类型的变量
*/
func main() {
	//var x = nil
	var y interface{} = nil
	//var z string = nil
	var n error = nil // error 类型，它是一种内置接口类型
	//fmt.Println(x, y, z, n)
	fmt.Println(y, n)

}
