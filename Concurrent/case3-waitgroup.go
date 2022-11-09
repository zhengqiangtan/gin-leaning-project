package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup
var urls = []string{
	"http://www.baidu.com",
	"https://www.csdn.net/",
	"http://www.wps.com",
}

// case3 : 多个goroutine同步机制
func case_waitgroup_1() {
	for _, url := range urls {
		wg.Add(1) // 设置需要等待的数目

		go func(url string) {
			defer wg.Done() // 请求结束后减1 等价于 wg.Add(-1)

			resp, err := http.Get(url)
			if err == nil {
				fmt.Println("get ", url, resp.Status)
			}
		}(url)
		wg.Wait() //等待所有请求结束
	}
}

func case_waitgroup_2() {
	// This WaitGroup is used to wait for all the
	// goroutines launched here to finish. Note: if a WaitGroup is
	// explicitly passed into functions, it should be done *by pointer*.
	var wg sync.WaitGroup

	// Launch several goroutines and increment the WaitGroup
	// counter for each.
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		// Avoid re-use of the same `i` value in each goroutine closure.
		// See [the FAQ](https://golang.org/doc/faq#closures_and_goroutines)
		// for more details.
		i := i

		// Wrap the worker call in a closure that makes sure to tell
		// the WaitGroup that this worker is done. This way the worker
		// itself does not have to be aware of the concurrency primitives
		// involved in its execution.
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	// Block until the WaitGroup counter goes back to 0;
	// all the workers notified they're done.
	wg.Wait()

	// Note that this approach has no straightforward way
	// to propagate errors from workers. For more
	// advanced use cases, consider using the
	// [errgroup package](https://pkg.go.dev/golang.org/x/sync/errgroup).
}

// This is the function we'll run in every goroutine.
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	// Sleep to simulate an expensive task.
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}
func main() {
	case_waitgroup_2()
}
