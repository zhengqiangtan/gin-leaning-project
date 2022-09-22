package main

import "fmt"

func case1_nil() {
	//var x = nil
	var y interface{} = nil
	//var z string = nil
	var n error = nil // error 类型，它是一种内置接口类型
	//fmt.Println(x, y, z, n)
	fmt.Println(y, n)
}

type Inter interface {
	Ping()
	Pang()
}

type St struct {
}

func (St) Ping() {
	fmt.Println("Ping....")
}

func (*St) Pang() {
	fmt.Println("Pang Pang")
}

func case2_nil() {

	var st *St = nil
	var it Inter = st
	fmt.Printf("%p\n", st) // 0x0
	fmt.Printf("%p\n", it) // 0x0

	if it != nil {
		it.Pang() // 这里居然可以打印出 Pang Pang

		//it.Ping() // panic: runtime error: invalid memory address or nil pointer dereference
		//GO的瑕疵：
		// 1. 空接口有两个字段，1个是实例类型 一个是指向绑定实例的指针，只有两个都为nil时，空接口才为nil

	}

}

/**
nil 只能赋值给指针、chan、func、interface、map 或 slice 类型的变量
*/
func main() {
	//case1_nil()
	case2_nil()
}
