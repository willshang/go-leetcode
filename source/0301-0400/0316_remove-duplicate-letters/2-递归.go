package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(removeDuplicateLetters("edaabca"))
}

// leetcode316_去除重复字母
func removeDuplicateLetters(s string) string {
	arr := [26]int{}
	pos := 0
	for i := 0; i < len(s); i++ {
		arr[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if s[i] < s[pos] {
			pos = i
		}
		arr[s[i]-'a']--
		if arr[s[i]-'a'] == 0 {
			break
		}
	}
	if len(s) == 0 {
		return ""
	}
	newStr := strings.ReplaceAll(s[pos+1:], string(s[pos]), "")
	return string(s[pos]) + removeDuplicateLetters(newStr)
}
