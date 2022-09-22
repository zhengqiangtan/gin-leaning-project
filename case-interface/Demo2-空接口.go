package main

import (
	"fmt"
)

// 空接口作为入参可以接受任何参数
func describe(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func main() {
	s := "Hello World"
	describe(s)
	i := 55
	describe(i)
	strt := struct {
		name string
	}{
		name: "Naveen R",
	}
	describe(strt)
}

//Type = string, value = Hello World
//Type = int, value = 55
//Type = struct { name string }, value = {Naveen R}
