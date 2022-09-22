package main

import (
	"fmt"
	"time"
)

//https://blog.csdn.net/foxlively/article/details/90292597
func main() {
	HelloMakeChanSize(1)
}

func HelloMakeChanSize(size int) {
	//size := 0
	c1 := make(chan int, size)
	go func() {
		for i := 0; i < 4; i++ {
			val := i*10 + 7
			fmt.Println(time.Now(), "<- ", val, "at", i)
			c1 <- i*10 + 7
		}
		c1 <- 0
	}()
	time.Sleep(time.Second * 3)
	fmt.Println("After Sleep")
	for val := range c1 {
		fmt.Println(time.Now(), "receive:", val)
		if val == 0 {
			break
		}
	}
}
