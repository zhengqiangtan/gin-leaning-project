package main

import (
	"fmt"
)

var (
	Ga int = 99
)

const (
	v int = 199
)

func GetGa() func() int {
	if Ga := 55; Ga < 60 {
		fmt.Println("GetGa if 中：", Ga)
	}
	for Ga := 2; ; {
		fmt.Println("GetGa循环中：", Ga)
		break
	}
	fmt.Println("GetGa函数中：", Ga)
	return func() int {
		Ga += 1
		return Ga
	}
}
func main() {
	Ga := "string"
	fmt.Println("main函数中：", Ga)
	b := GetGa() //闭包函数 return
	fmt.Println("main函数中：", b(), b(), b(), b())

	v := 1
	{
		v := 2
		fmt.Println(v)
		{
			v := 3
			fmt.Println(v)
		}
	}
	fmt.Println(v)
}

//main函数中： string
//GetGa if 中： 55
//GetGa循环中： 2
//GetGa函数中： 99
//main函数中： 100 101 102 103
//2
//3
//1
