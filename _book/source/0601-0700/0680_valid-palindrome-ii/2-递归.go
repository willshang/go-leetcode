package main

import "fmt"

func main() {
	//fmt.Println(validPalindrome("abca"))
	fmt.Println(validPalindrome("eccer"))
}

func validPalindrome(s string) bool {
	length := len(s)
	if length < 2 {
		return true
	}
	if s[0] == s[length-1] {
		return validPalindrome(s[1 : length-1])
	}
	return isPalindrome(s[0:length-1]) || isPalindrome(s[1:length])
}

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
