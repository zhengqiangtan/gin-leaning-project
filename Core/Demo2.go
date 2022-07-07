package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

/***
1、nil 测试用例
var nil Type // 类型必须是指针、通道、func、接口、映射或切片类型
/nil 可以被定义变量，但不建议，定义后便没有了原来的意义

2、 文件读取
*/
func ReadFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func main() {
	println("测试读取文件...")
	const filename = "D:\\go_workspace\\src\\gin-leaning-project\\Core\\test.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
	//fmt.Println(contents)
	//fmt.Println(err)

	// 测试读取2
	// 内建函数 make 用来为 slice，map 或 chan 类型分配内存和初始化一个对象(注意：只能用在这三种类型上)
	buf := make([]byte, 1024)
	f, _ := os.Open("D:\\go_workspace\\src\\gin-leaning-project\\Core\\test.txt")
	defer f.Close() // 延迟关闭文件流
	for {
		n, _ := f.Read(buf)
		if n == 0 {
			break

		}
		os.Stdout.Write(buf[:n])
	}
}
