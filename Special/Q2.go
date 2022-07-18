package main

import "fmt"

func hello() []string {
	return nil
}

func case1() {
	h := hello() // out:not nil  将hello()的返回值赋给h
	//h := hello // 是将 hello() 赋值给变量 h，而不是函数的返回值
	if h == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
}

func test2(num ...int) {
	num[0] = 18
}
func case2() {
	i := []int{5, 6, 7}
	test2(i...)
	fmt.Println(i[0])             // out:18
	fmt.Println(i[0], i[1], i[2]) // out:18,6,7
}

func case3() {
	s := make(map[string]int)
	delete(s, "h")
	fmt.Println(s["h"])
}

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct { //外部类型
	People //内部类型
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func case4() {
	t := Teacher{}
	t.ShowB() // teacher showB  内部定义的同名方法被屏蔽
	t.ShowA() // showA showB
}

func test5(i int) {
	fmt.Println(i)
}
func case5() {
	i := 5
	defer test5(i) // out:5 hello() 函数的参数在执行 defer 语句的时候会保存一份副本，在实际调用 hello() 函数时用
	i = i + 10
}

func incr(p *int) int {
	*p++
	return *p
}

func case6() {
	p := 1
	incr(&p)       // p 是 *int 类型的指针，这里指向的是变量 p 的地址，将该地址的值执行一个自增操作
	fmt.Println(p) //2
}

func add(args ...int) int {
	sum := 0
	for _, arg := range args {
		sum += arg
	}
	return sum
}
func case7() {
	a := add(1, 2)
	b := add(1, 2, 3)
	c := add([]int{1, 3, 7}...)
	fmt.Println(a, b, c)
}

func case8() {
	var s1 []int
	//var s2 = []int{} // 空切片和 nil 不相等，表示一个空的集合
	if s1 == nil { // nil 切片和 nil 相等 一般用来表示一个不存在的切片
		fmt.Println("yes nil")
	} else {
		fmt.Println("no nil")
	}
}

// 返回参数是匿名
func increaseA() int {
	var i int
	defer func() {
		i++
	}()
	return i
}

//返回参数是具名
func increaseB() (r int) {
	defer func() {
		r++
	}()
	return r
}
func case9() {
	fmt.Println(increaseA()) // 0
	fmt.Println(increaseB()) // 1
}

type A interface {
	ShowA() int
}

type B interface {
	ShowB() int
}

type Work struct {
	i int
}

func (w Work) ShowA() int {
	return w.i + 10
}

func (w Work) ShowB() int {
	return w.i + 20
}
func case10() {
	var a A = Work{3}
	s := a.(Work)
	fmt.Println(s.ShowA()) // 13
	fmt.Println(s.ShowB()) // 23

	//第二种情况
	c := Work{2}
	var t A = c            //t 的静态类型是 A
	fmt.Println(t.ShowA()) // 调用接口A的showA方法
}
func case11() {
	s := [3]int{1, 2, 3}
	a := s[:0]         //0-0,3-0
	b := s[:2]         //2-0,3-0
	c := s[1:2:cap(s)] //2-1,3-1
	fmt.Println(a, b, c)
	// 问：a,b,c的长度和容量分别是多少？
	// 答： 0 3、2 3、1 2
}

type Person struct {
	age int
}

func case12() {
	person := &Person{28}

	// 1.
	defer fmt.Println(person.age)
	// person.age 这一行代码跟之前含义是一样的，此时是将 28 当做 defer 函数的参数，会把 28 缓存在栈中，等到最后执行该 defer 语句的时候取出，即输出 28

	// 2.
	defer func(p *Person) {
		fmt.Println(p.age) // defer 缓存的是结构体 Person{28} 的地址，这个地址指向的结构体没有被改变，最后 defer 语句后面的函数执行的时候取出仍是 28；
	}(person)

	// 3.
	defer func() {
		fmt.Println(person.age) // 闭包引用，person 的值已经被改变，指向结构体 Person{29}
	}()

	person = &Person{29} //修改引用对象本身
	// out： 执行顺序为 3,2，1  =》 29 28 28

	//person = 29 //修改引用对象的成员 age
	// out： 执行顺序为 3,2，1  =》 29 29 28
}

func main() {

	//case1() //返回值问题
	//case2() //可变函数问题
	//case3() //map删除问题（ 删除 map 不存在的键值对时，不会报错，没有任何作用；获取不存在的减值对时，返回值类型对应的零值）
	//case4() //结构体嵌套问题,通过嵌套，内部类型的属性、方法，可以为外部类型所有，就好像是外部类型自己的一样。
	//case5() //defer问题
	//case6() // 指针问题 自增
	//case7() // 可变函数问题 同case2
	//case8() // nil 切片和空切片
	//case9()  // 具名和匿名函数问题 todo
	//case10() // 类型断言,Work实现了接口类型A ，调用自个的SHowA方法
	//case11() // 切片类型长度和容量换算
	case12() // 闭包测试&defer

}
