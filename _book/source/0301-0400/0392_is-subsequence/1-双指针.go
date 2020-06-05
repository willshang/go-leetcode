package main

import "fmt"

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc"))
}

// leetcode392_判断子序列
func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}
	i := 0
	j := 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == len(s)
}
