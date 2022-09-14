package main

import (
	"fmt"
	"math/rand"
	"net/url"
	"strconv"
	"time"
)

func case1_timeformat() {
	// 输出三十分钟前的日期
	startTime := time.Now().Add(-time.Minute * 31).Format("2022-01-01 15:04:05")
	fmt.Println(startTime)
	var escape = url.QueryEscape(startTime) // url编码
	fmt.Println(escape)                     // 6066-09-09+15%3A19%3A06
}

func case2_sleep() {
	// case2: 休眠测试
	fmt.Println("休眠开始：" + time.Now().GoString())
	<-time.After(time.Second * 3) // 休眠3s
	fmt.Println("休眠结束：" + time.Now().GoString())
}

func case3_ticker() {
	// case1: 每隔3s 打印一次
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		fmt.Println("归类转换......")
	}

}

func case4_variant() {
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		rand.Seed(time.Now().UnixNano())
		itoa := strconv.Itoa(rand.Intn(100))
		fmt.Println("每次调度参数都不同....." + itoa)
	}
}

func main() {
	//case1_timeformat()
	//case2_sleep()
	//case3_ticker()
	case4_variant()
}
