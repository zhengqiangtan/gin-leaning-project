package main

import (
	"fmt"
	"sync"
	"time"
)

var l sync.RWMutex

func main() {
	//go readAndRead()
	//
	//time.Sleep(1 * time.Second)
	//l.Lock()
	//fmt.Println("----------------- got lock")
	//l.Unlock()
	//
	//time.Sleep(5 * time.Second)

	test_mutex_2()
}

func readAndRead() {
	l.RLock()
	fmt.Println("----------------- got rlock")
	time.Sleep(10 * time.Second)
	fmt.Println("----------------- 10s passed")

	l.RLock()
	fmt.Println("----------------- got 2nd rlock")

	l.RUnlock()
	l.RUnlock()
}

/* shell 执行 `go run main.go` 的结果为：
   ----------------- got rlock
   ----------------- 10s passed
   fatal error: all goroutines are asleep - deadlock!
   ...
*/

func test_mutex_2() {
	var mutex sync.RWMutex
	arr := []int{1, 2, 3}
	go func() {
		fmt.Println("Try to lock writing operation.")
		mutex.Lock()
		fmt.Println("Writing operation is locked.")

		arr = append(arr, 4)

		fmt.Println("Try to unlock writing operation.")
		mutex.Unlock()
		fmt.Println("Writing operation is unlocked.")
	}()

	go func() {
		fmt.Println("Try to lock reading operation.")
		mutex.RLock()
		fmt.Println("The reading operation is locked.")

		fmt.Println("The len of arr is : ", len(arr))

		fmt.Println("Try to unlock reading operation.")
		mutex.RUnlock()
		fmt.Println("The reading operation is unlocked.")
	}()

	time.Sleep(time.Second * 2)
	return
}
