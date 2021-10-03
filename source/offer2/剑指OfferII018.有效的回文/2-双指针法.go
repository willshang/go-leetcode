package main

import (
	"fmt"
	"strings"
)

func main() {
	//str := "A man, a plan, a canal: Panama"
	str := "race a car"
	fmt.Println(isPalindrome(str))
}

// 剑指OfferII018.有效的回文
func isPalindrome(s string) bool {
	str := ""
	s = strings.ToLower(s)
	for _, value := range s {
		if (value >= '0' && value <= '9') || (value >= 'a' && value <= 'z') {
			str += string(value)
		}
	}
	if len(str) == 0 {
		return true
	}
	i := 0
	j := len(str) - 1
	for i <= j {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}
