package main

import "fmt"

// 线程池数量
const NUMBER = 10

// 任务结构
type tsk struct {
	begin int

	end int

	//只能进不能出的通道

	result chan<- int
}

// 任务执行逻辑
func (t *tsk) do() {

	sum := 0
	for i := t.begin; i <= t.end; i++ {
		sum += i
	}
	t.result <- sum

}

// 初始化线程池
func Inittsk(tskchan chan<- tsk, r chan int, p int) {

	qu := p / 10

	mod := p % 10

	hight := qu * 10

	for j := 0; j < qu; j++ {

		b := 10*j + 1

		e := 10 * (j + 1)

		task := tsk{
			begin:  b,
			end:    e,
			result: r,
		}

		tskchan <- task

	}

	if mod != 0 {
		task := tsk{
			begin:  hight + 1,
			end:    p,
			result: r,
		}
		tskchan <- task

	}
	close(tskchan)

}

// Distributetsk 分发任务
func Distributetsk(tskchan <-chan tsk, workers int, done chan struct{}) {

	for i := 0; i < workers; i++ {
		go Processtsk(tskchan, done)
	}

}

// Processtsk 执行任务
func Processtsk(tskchan <-chan tsk, done chan struct{}) {

	for t := range tskchan {
		t.do()
	}

	done <- struct{}{}
}

// CloseResult 关闭线程池
func CloseResult(done chan struct{}, resultchan chan int, workers int) {

	for i := 0; i < workers; i++ {
		<-done
	}

	close(done)
	close(resultchan)

}

// ProcessRes 获取任务执行结果
func ProcessRes(resultchan chan int) int {

	sum := 0
	for r := range resultchan {
		sum += r
	}
	return sum

}

func main() {
	workers := NUMBER

	// 以任务结构体为类型的通道，传递任务
	tskchan := make(chan tsk, 10)

	// 此通道用于任务处理结果的保存，传递tsk结果
	result := make(chan int, 10)

	// 同步线程池的通道，发送通知
	done := make(chan struct{}, 10)

	go Inittsk(tskchan, result, 100)

	Distributetsk(tskchan, workers, done)

	go CloseResult(done, result, workers)

	sum := ProcessRes(result)

	fmt.Println("sum = ", sum)

}

// 结果  5050
