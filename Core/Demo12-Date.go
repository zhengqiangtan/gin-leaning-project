package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now() // 获取当前时间

	fmt.Printf("当前时间:%v\n", t)
	fmt.Println("年", t.Year())
	fmt.Println("月", t.Month())
	fmt.Println("日", t.Day())
	fmt.Println("时", t.Hour())
	fmt.Println("分", t.Minute())
	fmt.Println("秒", t.Second())
}
