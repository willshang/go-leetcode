package main

import "fmt"

func main() {
	fmt.Println(shortestPalindrome("abcd"))
}

// leetcode214_最短回文串
func shortestPalindrome(s string) string {
	str := reverse(s)
	i := 0
	for i = 0; i < len(s); i++ {
		if str[i:] == s[:len(s)-i] {
			break
		}
	}
	return str[:i] + s
}

func reverse(s string) string {
	res := make([]byte, 0)
	for i := len(s) - 1; i >= 0; i-- {
		res = append(res, s[i])
	}
	return string(res)
}
