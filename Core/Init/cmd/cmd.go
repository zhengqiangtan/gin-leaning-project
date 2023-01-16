package main

// https://mp.weixin.qq.com/s/g4_HWOzL3NNFvyOHseLE6Q
import (
	"fmt"
	"gin-leaning-project/Core/Init/animals"
)

func main() {
	fmt.Println(animals.Random())
	fmt.Println(animals.Random())
	fmt.Println(animals.Random())
	fmt.Println(animals.Random())
	fmt.Println(animals.Random())

}
