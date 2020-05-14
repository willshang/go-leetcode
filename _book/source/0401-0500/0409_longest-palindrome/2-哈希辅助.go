package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("abcccdd"))
}

// leetcode409_最长回文串
func longestPalindrome(s string) int {
	ret := 0
	a := make(map[byte]int)
	for i := range s {
		a[s[i]]++
	}
	hasOdd := 0
	for i := range a {
		if a[i]%2 == 0 {
			ret += a[i]
		} else {
			ret += a[i] - 1
			hasOdd = 1
		}
	}
	return ret + hasOdd
}
