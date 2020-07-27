package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("cbbd"))
}

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	res := ""
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			str := s[i : j+1]
			if len(str) < len(res) && res != "" {
				continue
			}
			if judge(str) == true && len(res) < len(str) {
				res = str
			}
		}
	}
	return res
}

func judge(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
