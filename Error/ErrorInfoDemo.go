package main

import (
	"fmt"
	"log"
)

// 未完 todo
func main() {
	defer fmt.Println("panic退出前处理")
	log.Println("println日志")
	log.Panic("panic日志")
	log.Fatal("程序退出日志")
}
