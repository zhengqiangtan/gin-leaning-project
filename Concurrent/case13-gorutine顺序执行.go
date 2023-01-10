package main

import (
	"fmt"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

// 要使 A->B->C 顺序执行
func main() {
	A := make(chan struct{})
	B := make(chan struct{})
	C := make(chan struct{})
	waitGroup.Add(3)
	go func() {
		defer waitGroup.Done()
		testA(A, B)
	}()

	go func() {
		defer waitGroup.Done()
		testB(B, C)
	}()

	go func() {
		defer waitGroup.Done()
		testC(C)
	}()
	close(A)
	waitGroup.Wait()
	time.Sleep(10 * time.Second)
}

func testA(A chan struct{}, B chan struct{}) {
	<-A
	fmt.Println(1)
	close(B)
}

func testB(B chan struct{}, C chan struct{}) {
	<-B
	fmt.Println(2)
	close(C)
}

func testC(C chan struct{}) {
	<-C
	fmt.Println(3)
}
