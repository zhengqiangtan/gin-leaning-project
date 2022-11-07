package main

import (
	"fmt"
	"strings"
	"time"
)

type LogProgress struct {
	rc          chan string
	wc          chan string
	path        string
	influxDbDsn string
}

//读取模块
func (l *LogProgress) ReadFromFile() {

	l.rc <- "hello"
}

// 日志处理模块
func (l *LogProgress) process() {
	res := <-l.rc
	l.wc <- strings.ToUpper(res)
}

// 写入模块
func (l *LogProgress) WriteToInfluxDB() {
	fmt.Println(<-l.wc)
}

//封装
func main() {

	log := &LogProgress{
		rc:          make(chan string),
		wc:          make(chan string),
		path:        "./access.log",
		influxDbDsn: "user&pwd@localhost....",
	}
	go log.ReadFromFile()
	go log.process()
	go log.WriteToInfluxDB()

	time.Sleep(3 * time.Second)
	fmt.Println("hello")
}
