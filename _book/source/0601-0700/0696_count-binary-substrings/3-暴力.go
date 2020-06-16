package main

import (
	"fmt"
)

func main() {
	str := "00110011"
	fmt.Println(countBinarySubstrings(str))
}

var count int

func countBinarySubstrings(s string) int {
	count = 0
	for i := 1; i < len(s); i++ {
		if s[i-1] == '0' && s[i] == '1' {
			CountString(s, i-1, i)
		}
		if s[i-1] == '1' && s[i] == '0' {
			CountString(s, i-1, i)
		}
	}
	return count
}

func CountString(s string, left, right int) {
	leftStr := s[left]
	rightStr := s[right]
	for left >= 0 && right < len(s) && s[left] == leftStr && s[right] == rightStr {
		left--
		right++
		count++
	}
}
