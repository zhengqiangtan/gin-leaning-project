package main

/*
go协程 Go 协程（Goroutine）是与其他函数同时运行的函数。可以认为 Go 协程是轻量级的线程，由 Go 运行时来管理
每个 channel 都有一个类型。此类型是允许信道传输的数据类型。channel 是类型相关的，一个 channel 只能传递一种类型的值，这个类型需要在声明 channel 时指定

注意：
1、如果希望运行其他 Go 协程，Go 主协程必须继续运行着。如果 Go 主协程终止，则程序终止，于是其他 Go 协程也不会继续运行
*/
import (
	"fmt"
	"time"
)

func hi() {
	fmt.Println("Hello world goroutine")
}

func Producer(queue chan<- int) {
	for i := 0; i < 10; i++ {
		queue <- i //写入
		fmt.Println("create :", i)
	}
}

func Consumer(queue <-chan int) {
	for i := 0; i < 10; i++ {
		v := <-queue // 读出
		fmt.Println("receive:", v)
	}
}

/*
Goroutine 和 channel 共同使用
*/
func main() {
	go hi()
	time.Sleep(1 * time.Second) // Go 主协程休眠了 1 秒

	queue := make(chan int, 88)
	go Producer(queue)
	go Consumer(queue)
	time.Sleep(1 * time.Second) // Go 主协程休眠了 1 秒, 让协程有机会运行
}

//create : 0
//create : 1
//create : 2
//create : 3
//receive: 0
//receive: 1
//receive: 2
//receive: 3
//create : 4
//create : 5
//create : 6
//create : 7
//create : 8
//create : 9
//receive: 4
//receive: 5
//receive: 6
//receive: 7
//receive: 8
//receive: 9
