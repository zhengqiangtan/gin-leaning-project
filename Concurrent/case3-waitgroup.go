package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var urls = []string{
	"http://www.baidu.com",
	"https://www.csdn.net/",
	"http://www.wps.com",
}

// case3 : 多个goroutine同步机制
func main() {
	for _, url := range urls {
		wg.Add(1) // 设置需要等待的数目

		go func(url string) {
			defer wg.Done() // 请求结束后减1 等价于 wg.Add(-1)

			resp, err := http.Get(url)
			if err == nil {
				fmt.Println(resp.Status)
			}
		}(url)
		wg.Wait() //等待所有请求结束
	}
}
