package main

import "fmt"

/**
Go语言之make
make函数是Go的内置函数，它的作用是为slice、map或chan初始化并返回引用。make仅仅用于创建slice、map和channel，并返回它们的实例。
*/
func main() {
	//fixedArray()
	//println("------------------------------")
	//mutableArray()

	//var list = make([]FeedbackReplyModelDto,1, 1)
	//var list = make([]FeedbackReplyModelDto,0, 1)
	var list []FeedbackReplyModelDto

	// db select ...
	list = append(list, FeedbackReplyModelDto{
		Id:           "10",
		App:          "android",
		Modular:      "pub",
		QuestionDesc: "test",
		ContentEn:    "this is a test",
		ContentCn:    "这是个测试",
	})
	fmt.Println(len(list))

}

type FeedbackReplyModelDto struct {
	Id           string `json:"id"`
	App          string `json:"app"`
	Modular      string `json:"modular"`
	QuestionDesc string `json:"question_desc"`
	ContentEn    string `json:"content_en"`
	ContentCn    string `json:"content_cn"`
}

//new(T) 和 make(T,args) 是 Go 语言内建函数，用来分配内存，但适用的类型不同。
//new(T) 会为 T 类型的新值分配已置零的内存空间，并返回地址（指针），即类型为 *T的值。换句话说就是，返回一个指针，该指针指向新分配的、类型为 T 的零值。适用于值类型，如数组、结构体等。

//make(T,args) 返回初始化之后的 T 类型的值，这个值并不是 T 类型的零值，也不是指针 *T，是经过初始化之后的 T 的引用。make() 只适用于 slice、map 和 channel.

// case1 : 固定数组 [n]type
func fixedArray() {
	arr := [3]int{1, 2, 3}
	println(len(arr))

	arr2 := [...]int{1, 2, 3, 4}
	println(len(arr2))

	for i, e := range arr2 {
		println(i, e)
	}
}

//case2:可变数组 - slice
func mutableArray() {
	// 创建方式1
	arr3 := [...]int{1, 2, 3, 4}
	s1 := arr3[0:3]
	s2 := arr3[:3]
	s3 := arr3[2:3]
	println(s1, s2, s3)

	// 创建方式2 --make
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
