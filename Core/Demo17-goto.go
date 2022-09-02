package main

import (
	"errors"
	"fmt"
	"os"
)

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("被除数不能为0")
	}
	return a / b, nil
}

func main() {
	// 如果使用goto统一处理错误，变量需要提前统一定义
	var n1, n2 int
	var err error

	//var (
	//	n1  int
	//	n2  int
	//	err error
	//)

	n1, err = div(100, 10)
	if err != nil {
		goto HANDLE_ERROR_TAG
	}
	fmt.Println(n1)

	n2, err = div(100, 0)
	if err != nil {
		goto HANDLE_ERROR_TAG
	}
	fmt.Println(n2)

HANDLE_ERROR_TAG:
	fmt.Println("两个数相除错误：", err)
	os.Exit(1)
}
