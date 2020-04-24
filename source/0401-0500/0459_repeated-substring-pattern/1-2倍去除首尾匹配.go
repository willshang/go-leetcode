package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(repeatedSubstringPattern("axbaxb"))
	fmt.Println(repeatedSubstringPattern("abab"))
	fmt.Println(repeatedSubstringPattern("aabcbaa"))
	fmt.Println(repeatedSubstringPattern("aa"))
}

/*
	axbaxb => xbaxbaxbax
	abab => bababa
	aa => aa
*/
// leetcode459_重复的子字符串
func repeatedSubstringPattern(s string) bool {
	if len(s) == 0 {
		return false
	}

	size := len(s)
	ss := (s + s)[1 : size*2-1]
	return strings.Contains(ss, s)
}
