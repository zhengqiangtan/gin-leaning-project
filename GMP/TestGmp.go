package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

// GMP 分析
// 访问：go tool trace trace.out
func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	fmt.Println("hello world")
}
