package Test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func TestEmail(t *testing.T) {
	content := "I cant access my VIP account\nrobles.oswaldo@hotmail.com\nprogram says im a regular " +
		"customer but in my microsoft account i have de renweal of VIP until January 2023"

	bol := VerifyEmailFormat(content)
	fmt.Println(bol) // true

	isEmail(content)
	isEmail("123@567.com")
	isEmail("312@qq.com ")

}

// 正则匹配邮箱
func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(strings.TrimSpace(email))
}

// 校验内容是否为邮箱
func isEmail(email string) {
	result, _ := regexp.MatchString(`^([\w\.\_\-]{2,10})@(\w{1,}).([a-z]{2,4})$`, strings.TrimSpace(email))
	if result {
		println(`正确的邮箱`)
	} else {
		println(`不是邮箱`)
	}
}
