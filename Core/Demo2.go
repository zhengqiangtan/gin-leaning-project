package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

/***
nil 测试用例
var nil Type // 类型必须是指针、通道、func、接口、映射或切片类型
/nil 可以被定义变量，但不建议，定义后便没有了原来的意义
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
}
