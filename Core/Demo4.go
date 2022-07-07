package main

import "fmt"

/**
Go语言之make
make函数是Go的内置函数，它的作用是为slice、map或chan初始化并返回引用。make仅仅用于创建slice、map和channel，并返回它们的实例。
*/
func main() {
	// slice
	demo := make([]int, 10)
	fmt.Println("demo:", demo)
	// output: demo: [0 0 0 0 0 0 0 0 0 0]
	fmt.Println("len(demo):", len(demo))
	// output: len(demo): 10
	fmt.Println("cap(demo):", cap(demo))
	// output: cap(demo): 10

	// map
	demo2 := make(map[string]int) // key:string value:int
	demo2["zhangsan"] = 13
	demo2["lisi"] = 15
	fmt.Println(demo2) // map[lisi:15 zhangsan:13]

	// chan
	demo3 := make(chan int, 10)
	fmt.Println("demo3:", demo3)
	// output: demo3: 0xc000120000
	fmt.Println("len(demo3):", len(demo3))
	// output: len(demo3): 0
	fmt.Println("cap(demo3):", cap(demo3))
	// output: cap(demo3): 10

}

//new的作用是初始化一个指向类型的指针（*T）。使用new函数来分配空间，传递给new函数的是一个类型，不是一个值。返回的是指向这个新分配的零值的指针。
//
//make的作用是为slice、map或chan初始化并返回引用（T）。make仅仅用于创建slice、map和channel，并返回它们的实例。
