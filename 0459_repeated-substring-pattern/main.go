package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(repeatedSubstringPattern("abcdabcd"))
	fmt.Println(repeatedSubstringPattern("aabcbaa"))
}
func repeatedSubstringPattern(s string) bool {
	if len(s) == 0{
		return false
	}

	size := len(s)
	ss := (s+s)[1:size*2-1]

	return strings.Contains(ss,s)
}