package main

import "runtime"

func main() {
	println("GOMAXPROCS=", runtime.GOMAXPROCS(0)) // 16 默认本机核心数
	runtime.GOMAXPROCS(2)
	println("GOMAXPROCS=", runtime.GOMAXPROCS(0)) // 2

	println(runtime.Version())
	println(runtime.NumGoroutine())

}

//GOMAXPROCS= 16
//GOMAXPROCS= 2
//go1.17.3
//1
