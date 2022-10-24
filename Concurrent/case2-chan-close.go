package main

import (
	"fmt"
	"runtime"
)

// 2. goroutine运行结束后，写通道中的数据不会消失，他可以缓冲和适配两个goroutine处理速率不一致的情况
// 削峰 增大吞吐
func main() {
	c := make(chan struct{})
	ci := make(chan int, 100)
	go func(i chan struct{}, j chan int) {

		for i := 0; i < 10; i++ {
			ci <- i
		}
		close(ci)
		// 写通道
		c <- struct{}{} // 空写通道
	}(c, ci)

	println("num:", runtime.NumGoroutine()) // 2

	<-c //空等待 读通道c 通过它进行同步等待

	println("num:", runtime.NumGoroutine()) // 1

	for v := range ci {
		fmt.Println(v)
	}
}
