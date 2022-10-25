package main

import (
	"fmt"
	"time"
)

// FutureTask 在并发执行时用于传递参数和保存返回的结果
type FutureTask struct {
	// 用于传递参数
	args chan interface{}

	// 实际业务中可能还有很多其他的数据

	// 用于保存结果
	res chan interface{}
}

// execFutureTask 用于开启一个Future模式的线程
func execFutureTask(futureTask *FutureTask) {
	// 读取传入的参数
	fmt.Println("goroutine读取到的参数:", <-futureTask.args)

	// 这里可以执行具体的业务逻辑
	result := "执行完业务逻辑后得到的结果"

	// 将结果进行保存
	futureTask.res <- result
	defer close(futureTask.res)
	return
}

// 参考：https://blog.csdn.net/weixin_44829930/article/details/123606804
func main() {

	// 创建一个FutureTask并开启一个goroutine去执行
	futureTask := FutureTask{
		make(chan interface{}),
		make(chan interface{}),
	}
	go execFutureTask(&futureTask)

	// 向FutureTask传入参数，如果不传的话会死锁
	futureTask.args <- "main线程传入的参数"

	// 这里可以并行的去执行一些其他业务逻辑
	time.Sleep(1 * time.Second)

	// 读取线程执行的
	fmt.Println("主线程读取future模式下goroutine的结果:", <-futureTask.res)

}
