package main

import (
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

	bol := checkContainSubstrings("It's from my iPhone.", "It's from my iPhone", "Sent from my iPad Pro")
	print(bol)

	contains := strings.Contains("Kirimkan Dokumen Masalah", "Kirimkan Dokumen Masalah")
	println(contains)
}
