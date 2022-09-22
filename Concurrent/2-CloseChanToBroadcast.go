package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

func GenerateIntA(done chan struct{}) chan int {
	ch := make(chan int)
	go func() {
	Label:
		for {
			select {
			case ch <- rand.Int():
				// 对退出信号done的监听
			case <-done:
				break Label

			}
		}
		// 收到通知后关闭通道
		close(ch)
	}()
	return ch
}

//退出通知机制
//下游的消费者不需要随机数时显示通知生产者停止生产
func main() {
	done := make(chan struct{})
	ch := GenerateIntA(done)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// 告诉生产者停止生产
	close(done)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// 此时生产者已经退出
	fmt.Println("NumGo=", runtime.NumGoroutine())

}

//5577006791947779410
//8674665223082153551
//6129484611666145821
//0
//NumGo= 1
