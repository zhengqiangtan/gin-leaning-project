package main

import (
	"fmt"
	"sync"
)

type task struct {
	begin  int
	end    int
	result chan<- int
}

func (t *task) do() {
	sum := 0
	for i := t.begin; i <= t.end; i++ {
		sum += i
	}
	t.result <- sum
}

// 构建task并写入task通道
func InitTask(taskChan chan<- task, r chan int, p int) {
	qu := p / 10
	mod := p % 10
	high := qu * 10
	for j := 0; j < qu; j++ {
		b := 10*j + 1
		e := 10 * (j + 1)
		tsk := task{
			begin:  b,
			end:    e,
			result: r,
		}
		taskChan <- tsk
	}
	if mod != 0 {
		tsk := task{
			begin:  high + 1,
			end:    p,
			result: r,
		}
		taskChan <- tsk
	}
	close(taskChan)
}

// DistributeTask 读取taskchan 每个task启动一个worker goroutine 进行处理，并等待每个task运行完关闭结果通道
func DistributeTask(taskChan <-chan task, wait *sync.WaitGroup, result chan int) {
	for v := range taskChan {
		wait.Add(1)
		go ProcessTask(v, wait)
	}
	wait.Wait()
	close(result)
}

// ProcessTask 处理具体工作，并将结果发送到结果通道
func ProcessTask(t task, wait *sync.WaitGroup) {
	t.do()
	wait.Done()
}

// ProcessResult 读取结果通道，汇总结果
func ProcessResult(resultchan chan int) int {
	sum := 0
	for r := range resultchan {
		sum += r
	}
	return sum
}

/*
示例：计算100个自然数之和，将任务拆分为多个task进行处理

*/
func main() {
	// 1. 创建任务通道
	taskchan := make(chan task, 10)
	// 2.创建结果通道
	resultchan := make(chan int, 10)

	// 同步等待任务的执行
	wait := &sync.WaitGroup{}
	// 初始化task的goroutine 计算100个自然数之和
	go InitTask(taskchan, resultchan, 100)
	// 每个task启动一个goroutine去处理
	go DistributeTask(taskchan, wait, resultchan)

	// 通过结果通道获取结果进行处理
	sum := ProcessResult(resultchan)

	fmt.Println("sum=", sum) // sum= 5050

}
