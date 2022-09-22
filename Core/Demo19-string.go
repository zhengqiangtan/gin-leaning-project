package main

import (
	"log"
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

func main() {

	//bol := checkContainSubstrings("It's from my iPhone.", "It's from my iPhone", "Sent from my iPad Pro")
	//print(bol)
	//contains := strings.Contains("Kirimkan Dokumen Masalah", "Kirimkan Dokumen Masalah")
	//println(contains)

	str := "image/jpeg; name=\"image1.jpeg\""

	index := strings.Index(str, "\"")

	log.Println(index, " ", str[index+1:len(str)-1])

	var ct = ""
	split := strings.Split(str, ";")
	if len(split) > 1 {
		ct = split[0]
	}
	log.Println(ct)

}
