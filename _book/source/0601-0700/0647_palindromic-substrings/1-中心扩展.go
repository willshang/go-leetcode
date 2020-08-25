package main

import "fmt"

func main() {
	fmt.Println(countSubstrings("abc"))
}

// leetcode647_回文子串
func countSubstrings(s string) int {
	n := len(s)
	res := 0
	for i := 0; i < 2*n-1; i++ {
		left, right := i/2, i/2+i%2
		for ; 0 <= left && right < n && s[left] == s[right]; left, right = left-1, right+1 {
			res++
		}
	}
	return res
}
