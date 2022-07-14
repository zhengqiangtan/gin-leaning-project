package main

import "fmt"

/**
golang interface 定义与实现
参考：https://halfrost.com/go_interface/

*/
func case1() {
	var x interface{} = nil
	var y *int = nil

	interfaceIsNil(x) // empty
	interfaceIsNil(y) // non-empty
}

func interfaceIsNil(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}

// ----
//定义接口
type Programmer interface {
	WriteHelloWorld() string
	say()
}

// 定义结构体实现
type GoProgrammer struct {
}

func (g *GoProgrammer) say() {
	fmt.Println("hi how are you ?")
}

func (g GoProgrammer) WriteHelloWorld() string {
	return "hello world"
}

func case2() {
	var p Programmer
	p = new(GoProgrammer) // 通过接口去接收new的结构体会报错没有实现接口中的方法
	sayHi := p.WriteHelloWorld()
	//println(sayHi) 打印的很随机，不知道为什么
	fmt.Println(sayHi)
	p.say()
}

func main() {
	//case1() // nil判断
	case2() // 接口定义与实现示例
}
