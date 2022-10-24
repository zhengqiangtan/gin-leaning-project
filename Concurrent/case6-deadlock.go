package main

import (
	"fmt"
	"time"
)

func readChan(c chan int) {
	for {
		fmt.Println(<-c)
	}

}

func writeChan(c chan int) {
	c <- 'A'
}

func main() {

	c := make(chan int) //无缓冲通道
	//c <- 'A'
	//go readChan(c) // fatal error: all goroutines are asleep - deadlock!
	// 原因：main函数不能继续向下执行了

	//改进方式
	go writeChan(c)
	go readChan(c)

	time.Sleep(time.Millisecond) // 必须等待一会

}
