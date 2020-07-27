package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("cbbd"))
}

// leetcode5_最长回文子串
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	start := 0
	end := 0
	for i := 0; i < len(s); i++ {
		left1, right1 := find(s, i, i)
		left2, right2 := find(s, i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

func find(s string, left, right int) (int, int) {
	for ; 0 <= left && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}
