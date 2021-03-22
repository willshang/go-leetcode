package main

import "fmt"

func main() {
	fmt.Println(checkPalindromeFormation("ulacfd", "jizalu"))
	fmt.Println(checkPalindromeFormation("aabcc", "ccdaa"))
}

func checkPalindromeFormation(a string, b string) bool {
	n := len(a)
	left := 0
	for left <= n/2 {
		if a[left] == b[n-1-left] {
			left++
		} else {
			break
		}
	}
	// 超过一半直接返回true
	if left >= n/2 {
		return true
	}
	// 判断剩下的
	if isPalindrome() || isPalindrome() {
		return true
	}
	return false
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
