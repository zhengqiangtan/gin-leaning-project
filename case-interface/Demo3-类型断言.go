package main

import (
	"fmt"
)

func assert(i interface{}) {
	s, ok := i.(int) //get the underlying int value from i
	fmt.Println(s, ok)
}

func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("I am a string and my value is %s\n", i.(string))
	case int:
		fmt.Printf("I am an int and my value is %d\n", i.(int))
	default:
		fmt.Printf("Unknown type\n")
	}
}

type Describer interface {
	Describe()
}
type Person struct {
	name string
	age  int
}

func (p Person) Describe() {
	fmt.Printf("%s is %d years old", p.name, p.age)
}
func findType2(i interface{}) {
	switch v := i.(type) {
	case Describer:
		v.Describe()
	default:
		fmt.Printf("unknown type\n")
	}
}

// 类型断言用于提取接口的底层值（Underlying Value）
func main() {
	//1 类型判断
	var s interface{} = 56
	assert(s)
	var a interface{} = "test"
	assert(a)
	// 2 类型选择
	findType("name is liming")
	findType(123)
	findType(123.234)

	// 3 将一个类型和接口相比较。如果一个类型实现了接口，那么该类型与其实现的接口就可以互相比较。
	p := Person{
		name: "Naveen R",
		age:  25,
	}
	findType2(p)

}
