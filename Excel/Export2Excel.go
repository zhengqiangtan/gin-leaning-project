package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	const filename = "C:\\Users\\HNK7WC3\\Downloads\\raw.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}
