package main

import (
	"fmt"
)

func main() {
	str := "00110011"
	fmt.Println(countBinarySubstrings(str))
}

// leetcode696_计数二进制子串
func countBinarySubstrings(s string) int {
	res := 0
	cur := 1
	pre := 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			cur++
		} else {
			if pre > cur {
				res = res + cur
			} else {
				res = res + pre
			}
			pre = cur
			cur = 1
		}
	}
	if pre > cur {
		return res + cur
	}
	return res + pre
}
