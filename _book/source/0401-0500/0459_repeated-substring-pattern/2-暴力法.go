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
	if len(s) == 0 {
		return false
	}
	size := len(s)
	for i := 1; i < size; i++ {
		if size%i == 0 {
			count := size / i
			if strings.Repeat(s[0:i], count) == s {
				return true
			}
		}
	}
	return false
}
