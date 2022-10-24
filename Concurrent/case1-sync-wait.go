package main

import "runtime"

// 使用无缓冲的通道来实现goroutines之间的同步等待
func main() {
	c := make(chan struct{})

	go func(i chan struct{}) {
		sum := 0
		for i := 0; i < 10000; i++ {
			sum += i
		}
		println(sum)
		c <- struct{}{} // 写通道
	}(c)

	println("num:", runtime.NumGoroutine())

	<-c // 读通道c 通过它进行同步等待sum计算完成
}

//num: 2
//49995000
