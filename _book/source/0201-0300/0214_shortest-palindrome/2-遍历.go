package main

import "fmt"

func main() {
	fmt.Println(shortestPalindrome("abcd"))
}

func shortestPalindrome(s string) string {
	i := len(s)
	for {
		if isPalindrome(s[:i]) == true {
			break
		}
		i--
	}
	res := s
	for j := i; j < len(s); j++ {
		res = string(s[j]) + res
	}
	return res
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
