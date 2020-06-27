package main

import "fmt"

func main() {
	fmt.Println(balancedStringSplit("RLRRLLRLRL"))
}

// leetcode1221_分割平衡字符串
func balancedStringSplit(s string) int {
	res := 0
	if len(s) == 0 {
		return res
	}
	left := 0
	right := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'L' {
			left++
		} else {
			right++
		}
		if left == right {
			res++
		}
	}
	return res
}
