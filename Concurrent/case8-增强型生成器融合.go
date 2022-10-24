package main

import (
	"fmt"
	"math/rand"
)

func gen_int_11(done chan struct{}) chan int {
	ch := make(chan int, 5)
	go func() {
	Label:
		for {
			t := rand.Int()
			fmt.Println("gen_int_11 gen data: ", t)
			select {
			case ch <- t:
			case <-done:
				break Label

			}
		}
		close(ch)
	}()
	return ch
}

func gen_int_22(done chan struct{}) chan int {
	ch := make(chan int, 10)
	go func() {
	Label:
		for {
			t := rand.Int()
			fmt.Println("gen_int_22 gen data: ", t)
			select {
			case ch <- t:
			case <-done:
				break Label

			}
		}
		close(ch)
	}()
	return ch
}

func gen_int_00(done chan struct{}) chan int {
	ch := make(chan int)
	send := make(chan struct{})

	go func() {
	Label:
		for {
			select {
			case ch <- <-gen_int_11(send):
			case ch <- <-gen_int_22(send):
			case <-done:
				send <- struct{}{}
				send <- struct{}{}
				break Label

			}
		}
		close(ch)
	}()
	return ch
}

func main() {
	// 创建一个退出信号的chan
	done := make(chan struct{})
	ch := gen_int_00(done) // 启动生成器

	// 获取生成器生产的随机值
	for i := 0; i < 5; i++ {
		fmt.Println("消费数据：", <-ch)
	}

	// 通知生产者停止生成
	done <- struct{}{}

	fmt.Println("已经停止生产")
}
