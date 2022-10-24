package main

import (
	"fmt"
	"math/rand"
)

func gen_int_1() chan int {
	ch := make(chan int, 10)
	go func() {
		for {
			ch <- rand.Int()
		}
	}()
	return ch
}

func gen_int_2() chan int {
	ch := make(chan int, 10)
	go func() {
		for {
			ch <- rand.Int()
		}
	}()
	return ch
}

func gen_int() chan int {
	ch := make(chan int, 20)

	go func() {
		for true {
			select { //使用select扇入技术增加随机性
			case ch <- <-gen_int_1():
			case ch <- <-gen_int_2():

			}
		}
	}()
	return ch
}

func main() {
	ch := gen_int()
	for i := 0; i < 100; i++ {
		fmt.Println(<-ch)
	}
}
