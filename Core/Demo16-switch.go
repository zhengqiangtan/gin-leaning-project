package main

import "fmt"

func simpleSwitch() {
	var evaluateValue = "good"
	switch evaluateValue {
	case "good":
		fmt.Println("this is good")
	case "bad":
		fmt.Println("this is bad")
	}
}

func exprSwitch() {
	var score int
	fmt.Scan(&score)
	switch {
	case score < 60:
		fmt.Println("不及格")
	case score >= 60 && score < 80:
		fmt.Println("不错")
	case score >= 80:
		fmt.Println("优秀")
	default:
		panic("wrong number")
	}
}
func main() {
	simpleSwitch() // 简单选择
	exprSwitch()   // 可以输入表达式判断

}
