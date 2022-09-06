package main

import (
	"fmt"
	"time"
)

func main() {

	//// case1: 每隔3s 打印一次
	//ticker := time.NewTicker(3 * time.Second)
	//defer ticker.Stop()
	//
	//for range ticker.C {
	//	fmt.Println("归类转换......")
	//}

	// case2: 休眠测试
	fmt.Println("休眠开始：" + time.Now().GoString())
	<-time.After(time.Second * 3) // 休眠3s
	fmt.Println("休眠结束：" + time.Now().GoString())

}
