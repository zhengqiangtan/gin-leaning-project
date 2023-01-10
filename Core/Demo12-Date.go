package main

import (
	"fmt"
	"github.com/jinzhu/now"
	"time"
)

func main() {
	//testTime01()
	//testTime02()

	day := time.Now().Day()
	fmt.Println(day)

}

func case1_time_type() {
	t := time.Now() // 获取当前时间

	fmt.Printf("当前时间:%v\n", t)
	fmt.Println("年", t.Year())
	fmt.Println("月", t.Month())
	fmt.Println("日", t.Day())
	fmt.Println("时", t.Hour())
	fmt.Println("分", t.Minute())
	fmt.Println("秒", t.Second())
}

func testTime01() {
	localTime := LocalTime()
	t := time.Now().In(localTime)
	day := now.With(t).EndOfDay()
	print(day.Second())
}

// +8 时区
func LocalTime() *time.Location {
	return time.FixedZone("UTC-8", +8*60*60)
}

func testTime02() {
	nowTime := time.Now()
	fmt.Println("nowTime.Location", nowTime.Location())
	zoneName, _ := nowTime.Zone()
	fmt.Println("nowTime.Zone", zoneName)
	fmt.Println("nowTime.IsZero", nowTime.IsZero())
	fmt.Println("nowTime.Location", nowTime.Local())
	fmt.Println("nowTime.UTC", nowTime.UTC())
	fmt.Println("nowTime.In", nowTime.In(nowTime.Location()))
	fmt.Println("nowTime.Unix", nowTime.Unix())
	fmt.Println("nowTime.UnixNano", nowTime.UnixNano())
	fmt.Println("nowTime.Equal", nowTime.Equal(nowTime))
	fmt.Println("nowTime.Before", nowTime.Before(nowTime))
	fmt.Println("nowTime.After", nowTime.After(nowTime))
	fmt.Println(nowTime.Date())
	fmt.Println(nowTime.Clock())
	fmt.Println("nowTime.Year", nowTime.Year())
	fmt.Println("nowTime.YearDay", nowTime.YearDay())
	fmt.Println("nowTime.Month", nowTime.Month())
	fmt.Println(nowTime.ISOWeek())
	fmt.Println("nowTime.Day", nowTime.Day())
	fmt.Println("nowTime.Weekday", nowTime.Weekday())
	fmt.Println("nowTime.Hour", nowTime.Hour())
	fmt.Println("nowTime.Minute", nowTime.Minute())
	fmt.Println("nowTime.Second", nowTime.Second())
	fmt.Println("nowTime.Nanosecond", nowTime.Nanosecond())
	fmt.Println("nowTime.Add", nowTime.Add(time.Hour*3))
	fmt.Println("nowTime.AddDate", nowTime.AddDate(2, 3, 4))
	fmt.Println("nowTime.Sub", nowTime.Sub(nowTime))
	fmt.Println("nowTime.String", nowTime.String())
	fmt.Println("nowTime.Format", nowTime.Format("2006-01-02 15:04:05"))
	fmt.Println("nowTime.Format", nowTime.Format("2006-01-02 15-04-05"))
	fmt.Println("nowTime.Format", nowTime.Format("2006-01-02"))
}
