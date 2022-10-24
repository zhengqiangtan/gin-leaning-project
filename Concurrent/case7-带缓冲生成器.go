package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ch := GenerateInt()
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func GenerateInt() chan int {
	ch := make(chan int, 10)
	go func() {
		for {
			ch <- rand.Int()
		}
	}()
	return ch
}
