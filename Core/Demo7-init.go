package main

import (
	"fmt"
	"strconv"
)

/**
 Golang init函数释义：
A. 一个包中，可以包含多个 init 函数；
B. 程序编译时，先执行依赖包的 init 函数，再执行 main 包内的 init 函数；

init() 函数是用于程序执行前做包的初始化的函数，比如初始化包里的变量等;
一个包可以出线多个 init() 函数,一个源文件也可以包含多个 init() 函数；
同一个包中多个 init() 函数的执行顺序没有明确定义，但是不同包的init函数是根据包导入的依赖关系决定的（看下图）;
init() 函数在代码中不能被显示调用、不能被引用（赋值给函数变量），否则出现编译错误;
一个包被引用多次，如 A import B,C import B,A import C，B 被引用多次，但 B 包只会初始化一次；
引入包，不可出现死循坏。即 A import B,B import A，这种情况编译失败；

*/
func hello() []string {
	return nil
}

func test_hello() {
	h := hello()
	if h == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
}

func main() {

	uint64_test()
}

func uint64_test() {
	i := uint64(123)
	s := strconv.FormatUint(i, 10)
	fmt.Println(i, s)
}
