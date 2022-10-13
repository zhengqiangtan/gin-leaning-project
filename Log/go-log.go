package main

import "log"

// Print系列(Print|Printf|Println）、Fatal系列（Fatal|Fatalf|Fatalln）、和Panic系列（Panic|Panicf|Panicln）
func main() {
	// 设置标准logger的输出配置
	log.SetFlags(log.Llongfile | log.Ltime | log.Ldate)
	log.SetPrefix("[TEST]")

	log.Println("这是一条很普通的日志。")
	v := "很普通的"
	log.Printf("这是一条%s日志。\n", v)
	log.Fatalln("这是一条会触发fatal的日志。")
	log.Panicln("这是一条会触发panic的日志。")
}
