package main

import (
	"fmt"
	"strings"
	"time"
)

// 使用接口优化

type Reader interface {
	Read(rc chan string)
}
type Writer interface {
	Write(rc chan string)
}

//整个功能模块的封装-结构通体
type LogProgress struct {
	rc    chan string
	wc    chan string
	read  Reader
	write Writer
}

type ReadFromFile struct {
	path string
}
type WriteToInfluxDB struct {
	influxDbDsn string
}

//实现接口-读取
func (r *ReadFromFile) Read(rc chan string) {
	rc <- "hello"
}

//实现接口-写入
func (w *WriteToInfluxDB) Write(wc chan string) {
	fmt.Println(<-wc)
}

// 日志处理模块
func (l *LogProgress) process() {
	res := <-l.rc
	l.wc <- strings.ToUpper(res)
}

func main() {

	r := &ReadFromFile{
		path: "./access.log",
	}

	w := &WriteToInfluxDB{
		influxDbDsn: "user:pwd@localhost",
	}

	log := &LogProgress{
		rc:    make(chan string),
		wc:    make(chan string),
		read:  r,
		write: w,
	}

	go log.read.Read(log.rc)
	go log.process()
	go log.write.Write(log.wc)

	//主线程等待
	time.Sleep(1 * time.Second)

}
