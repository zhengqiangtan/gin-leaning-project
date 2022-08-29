package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		fmt.Println("归类转换......")
	}

}
