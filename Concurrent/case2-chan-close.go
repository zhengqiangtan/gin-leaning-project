package main

import "runtime"

func main() {
	c := make(chan struct{})
	go func(i chan struct{}) {
		sum := 0
		for i := 0; i < 10000; i++ {
			sum += i
		}
		println(sum)
		c <- struct{}{} // 空写通道
	}(c)

	println("num:", runtime.NumGoroutine())

	<-c //空等待 读通道c 通过它进行同步等待sum计算完成
}
