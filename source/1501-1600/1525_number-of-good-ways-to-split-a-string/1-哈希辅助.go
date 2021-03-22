package main

import "fmt"

func main() {
	fmt.Println(numSplits("abcd"))
}

// leetcode1525_字符串的好分割数目
func numSplits(s string) int {
	left := make(map[byte]int)
	right := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		left[s[i]]++
	}
	res := 0
	for i := 0; i < len(s); i++ {
		left[s[i]]--
		right[s[i]]++
		if left[s[i]] == 0 {
			delete(left, s[i])
		}
		if len(left) == len(right) {
			res++
		}
	}
	return res
}
