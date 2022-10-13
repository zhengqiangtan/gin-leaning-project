package main

import (
	"fmt"
	"github.com/wxnacy/wgo/arrays"
	"strings"
)

func checkContainSubstrings(str string, subs ...string) bool {
	isCompleteMatch := false
	for _, sub := range subs {
		if strings.Contains(str, sub) {
			isCompleteMatch = true
		}
	}
	return isCompleteMatch
}

func regexFuc() {
	var arr = make([]string, 5)
	arr = append(arr, "aa", "bb", "cc")

	var str = "aa"
	s := strings.Replace(str, ";", "|", -1)
	fmt.Println(s)

}

func stringContains() {
	sa := []string{"created_at", "feedback_id", "content", "r", "t"}

	bol := arrays.ContainsString(sa, "created_at") // 0
	fmt.Println(bol)
}

func main() {
	//regexFuc()
	stringContains() //字符串在数组中包含判断

}
