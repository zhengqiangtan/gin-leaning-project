package main

/*
 只要监听的通道中有一个是可读或可写的，select就不会阻塞，而是进入处理就绪通道的分支流程,
 如果监听的通道中有多个可读或者可写，则select随机选择一个处理

*/
func main() {
	ch := make(chan int, 1)

	go func(chan int) {
		for {
			select {
			// 0 或1的写入是随机的
			case ch <- 0:
			case ch <- 1:
			}
		}
	}(ch)

	for i := 0; i < 10; i++ {
		println(<-ch)
	}
}
