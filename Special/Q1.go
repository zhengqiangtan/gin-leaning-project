package main

import (
	"fmt"
)

/**
1、range常犯的错误：
for range 循环的时候会创建每个元素的副本，而不是元素的引用，所以 m[key] = &val 取的都是变量 val 的地址，
所以最后 map 中的所有元素的值都是变量 val 的地址，因为最后 val 被赋值为3，所有输出都是3

2、defer释义:
defer 的执行顺序是后进先出。当出现 panic 语句的时候，会先按照 defer 的后进先出的顺序执行，最后才会执行panic
*/

func rangeTest() {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
		//m[key] = &val
		// 修改
		value := val
		m[key] = &value

	}

	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
	// bad out
	//1 -> 3
	//2 -> 3
	//3 -> 3
	//0 -> 3
}

func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
	panic("触发异常")
}

//打印后
//打印中
//打印前
//panic: 触发异常

func appendTest1() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

func appendTest2() {
	s := make([]int, 0)
	s = append(s, 1, 2, 3) // 注意： append(s1, s2) append() 的第二个参数不能直接使用 slice，应该直接跟元素append(s1,1,2,3)
	fmt.Println(s)         // [1,2,3]
}

//缺陷：第二个参数没有返回值 ，加上a错误消失
// 在函数有多个返回值时，只要有一个返回值有命名，其他的也必须命名
func funcMui(x, y int) (sum int, a error) {
	return x + y, nil
}

func testNewIntArray() {
	//list := new([]int) // new([]int) 之后的 list 是一个 *[]int 类型的指针，不能对指针执行 append 操作
	//list = append(list,1)
	//fmt.Println(list)
}

var (
	size = 1024
	//size := 1024 // bad
	max_size = size * 2
)

func structCompareTest() {
	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	//sm1 := struct {
	//	age int
	//	m   map[string]string
	//}{age: 11, m: map[string]string{"a": "1"}}
	//sm2 := struct {
	//	age int
	//	m   map[string]string
	//}{age: 11, m: map[string]string{"a": "1"}}
	//
	//if sm1 == sm2 {
	//	fmt.Println("sm1 == sm2")
	//}
	// 那什么是可比较的呢，常见的有 bool、数值型、字符、指针、数组等，
	// 像slice、map、函数等是不能比较的

}

type MyInt1 int   //基于类型 int 创建了新类型 MyInt1
type MyInt2 = int // 创建了 int 的类型别名 MyInt2

func main() {
	//rangeTest()
	//defer_call()
	//appendTest1()
	//appendTest2()

	// 短变量声明 必须在函数内部进行
	//fmt.Println(size, max_size)

	//structCompareTest()

	// 类型别名与类型定义的区别
	var i int = 0
	//var i1 MyInt1 = i // bad use
	var i2 MyInt2 = i // good use  本质还是int类型
	//fmt.Println(i1, i2)

	var a MyInt1 = MyInt1(i) // 使用强制类型转化
	fmt.Println(a, i2)

}
