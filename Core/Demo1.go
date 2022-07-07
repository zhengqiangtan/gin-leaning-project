package main

import (
	"fmt"
	"unsafe"
)

/**
组合用例，在定义的cat 中使用Animal结构体
通过定义不同字段和方法的结构体，抽象组合不同的结构体，这大概便是Go语言中对面向对象编程了
*/
type Animal struct {
	Name   string  //名称
	Color  string  //颜色
	Height float32 //身高
	Weight float32 //体重
	Age    int     //年龄
}

func (a Animal) Run() {
	fmt.Println(a.Name + "正在跑")
}
func (a Animal) Eat() {
	fmt.Println(a.Name + "正在吃")
}

type Cat struct {
	a    Animal
	miao string
}
type Lion struct {
	Animal // 匿名字段类型
}

func main() {
	var c = Cat{
		a: Animal{
			Name:   "猫猫",
			Color:  "橙色",
			Weight: 10,
			Height: 30,
			Age:    1,
		},
		miao: "cat会喵喵叫",
	}
	res := fmt.Sprintf("this is %v", c.a.Name)
	fmt.Println(res)
	fmt.Println(c.miao)
	c.a.Run() // 非常繁琐，可以使用匿名字段来解决,如小狮子定义

	var lion = Lion{
		Animal{
			Name:  "小狮子",
			Color: "黄色",
		},
	}
	lion.Run()
	fmt.Println(lion.Color)

	// 查看结构体占了多少内存
	size := unsafe.Sizeof(Animal{})
	fmt.Println(size) // 48

	// 对象调用测试
	animal := &Animal{Name: "旺财", Age: 8, Color: "黄色"}
	fmt.Println(animal) //&{旺财 黄色 0 0 8}

}
