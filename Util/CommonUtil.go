package Util

import (
	"regexp"
)

//func main() {
//	fmt.Println(VerifyEmailFormat("12345@126.com")) //true
//	fmt.Println(VerifyEmailFormat("12345126.com"))  //false
//}

func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}
