package main

// 参考：https://youwu.today/skill/backend/using-enum-in-golang/
import "fmt"

const (
	Running int = iota // 常量计数器，iota，它使用在声明常量时为常量连续赋值
	Pending
	Stopped
)

func main() {
	fmt.Println("State running: ", Running)
	fmt.Println("State pending: ", Pending)
	fmt.Println("State Stoped: ", Stopped)
}
