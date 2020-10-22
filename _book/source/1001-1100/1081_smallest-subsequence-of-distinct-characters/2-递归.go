package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(smallestSubsequence("bca"))
}

func smallestSubsequence(text string) string {
	arr := [26]int{}
	pos := 0
	for i := 0; i < len(text); i++ {
		arr[text[i]-'a']++
	}
	for i := 0; i < len(text); i++ {
		if text[i] < text[pos] {
			pos = i
		}
		arr[text[i]-'a']--
		if arr[text[i]-'a'] == 0 {
			break
		}
	}
	if len(text) == 0 {
		return ""
	}
	newStr := strings.ReplaceAll(text[pos+1:], string(text[pos]), "")
	return string(text[pos]) + smallestSubsequence(newStr)
}
