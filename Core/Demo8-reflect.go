package main

import (
	"fmt"
	"reflect"
)

/**
golang 反射示例
https://halfrost.com/go_reflection/
*/
func main() {
	name := "Do"
	t := &T{}
	reflect.ValueOf(t).MethodByName(name).Call(nil)
}

type T struct {
}

func (t *T) Do() {
	fmt.Println("hello")
}
