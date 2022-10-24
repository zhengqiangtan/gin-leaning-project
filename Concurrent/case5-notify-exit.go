package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

// 生成随机数

func GenerateIntA(done chan struct{}) chan int {
	ch := make(chan int)
	go func() {
	Label:
		for {
			i := rand.Int()
			fmt.Println("product-->", i)
			select {
			case ch <- i:
			case <-done: // 对退出信号done的监听
				break Label

			}
		}
		// 收到通知后关闭通道
		close(ch)
	}()
	return ch
}

//退出通知机制示例：
//功能：下游的消费者不需要随机数时显示通知生产者停止生产
func main() {
	done := make(chan struct{})
	ch := GenerateIntA(done)
	fmt.Println("consumer-->", <-ch)
	fmt.Println("consumer-->", <-ch)

	fmt.Println("NumGo=", runtime.NumGoroutine())
	// 告诉生产者停止生产
	fmt.Println("通知：生产者停止生产")
	close(done)

	fmt.Println("consumer-->", <-ch)
	fmt.Println("consumer-->", <-ch)

	// 此时生产者已经退出
	fmt.Println("NumGo=", runtime.NumGoroutine())

}

//product--> 5577006791947779410
//product--> 8674665223082153551
//consumer--> 5577006791947779410
//consumer--> 8674665223082153551
//NumGo= 2
//通知：生产者停止生产
//product--> 6129484611666145821
//consumer--> 0
//consumer--> 0
//NumGo= 1
