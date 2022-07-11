package main

import "fmt"

/**
iota是golang语言的常量计数器,只能在常量的表达式中使用
iota在const关键字出现时将被重置为0(const内部的第一行之前)，const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)
*/
type Stereotype int

const (
	TypicalNoob           Stereotype = iota // 0
	TypicalHipster                          // 1
	TypicalUnixWizard                       // 2
	TypicalStartupFounder                   // 3
)

// case2: 日期
type Weekday int

const (
	Sunday Weekday = iota // 0
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// case3: 使用下划线跳过不想要的值。
type AudioOutput int

const (
	OutMute   AudioOutput = iota // 0
	OutMono                      // 1
	OutStereo                    // 2
	_
	_
	OutSurround // 5
)

// case4: 定义数量级
type ByteSize float64

const (
	_           = iota             // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota) // 1 << (10*1)
	MB                             // 1 << (10*2)
	GB                             // 1 << (10*3)
	TB                             // 1 << (10*4)
	PB                             // 1 << (10*5)
	EB                             // 1 << (10*6)
	ZB                             // 1 << (10*7)
	YB                             // 1 << (10*8)
)

// case5: 中间插队的情况
const (
	i = iota
	j = 3.14
	k = iota
	l
)

// out: 0 3.14 2 3

func main() {

	fmt.Println(TypicalNoob, TypicalHipster, TypicalUnixWizard, TypicalStartupFounder)
	fmt.Println(KB, MB, GB, TB)
	fmt.Println(i, j, k, l)
}
