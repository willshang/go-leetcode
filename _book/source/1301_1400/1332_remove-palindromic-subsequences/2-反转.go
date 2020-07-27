package main

import "fmt"

func main() {
	fmt.Println(removePalindromeSub("ababa"))
}

func removePalindromeSub(s string) int {
	if len(s) == 0 {
		return 0
	}
	temp := ""
	for i := len(s) - 1; i >= 0; i-- {
		temp = temp + string(s[i])
	}
	if temp == s {
		return 1
	}
	return 2
}
