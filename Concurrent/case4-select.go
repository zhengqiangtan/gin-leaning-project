package main

import (
	"fmt"
	"time"
)

/*
 只要监听的通道中有一个是可读或可写的，select就不会阻塞，而是进入处理就绪通道的分支流程,
 如果监听的通道中有多个可读或者可写，则select随机选择一个处理

*/
func main() {
	//c := make(chan string, 1)
	////c <- "aa"
	//selectDemo(c)

	case_select_timeout()
}

func select_case1() {
	ch := make(chan int, 1)

	go func(chan int) {
		for {
			select {
			// 0 或1的写入是随机的
			case ch <- 0:
			case ch <- 1:
			}
		}
	}(ch)

	for i := 0; i < 10; i++ {
		println(<-ch)
	}
}

func selectDemo(c chan string) {
	recv := ""
	send := "Hello"
	select {
	case recv = <-c:
		fmt.Printf("Received %s\n", recv)
	case c <- send:
		fmt.Printf("Sent %s\n", send)
	}
}

func case_select_timeout() {
	ch := make(chan int)
	quit := make(chan bool)

	//新开一个协程
	go func() {
		for {
			select {
			case num := <-ch:
				fmt.Println("num = ", num)
			case <-time.After(3 * time.Second):
				fmt.Println("超时")
				quit <- true
			}
		}

	}() //别忘了()

	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Second)
	}

	<-quit
	fmt.Println("程序结束")
}
