package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("abcccdd"))
}

func longestPalindrome(s string) int {
	ret := 0
	a := [123]int{}
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
