package main

import "fmt"

/**
golang interface
参考：https://halfrost.com/go_interface/

*/
func main() {
	var x interface{} = nil
	var y *int = nil

	interfaceIsNil(x) // empty
	interfaceIsNil(y) // non-empty

}

func interfaceIsNil(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}
