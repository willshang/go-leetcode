package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(removeDuplicates("deeedbbcccbdaa", 3))
}

// leetcode1209_删除字符串中的所有相邻重复项II
func removeDuplicates(s string, k int) string {
	if len(s) < k {
		return s
	}
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		stack = append(stack, s[i])
		if len(stack) >= k {
			a := string(stack[len(stack)-k:])
			b := strings.Repeat(string(s[i]), k)
			if a == b {
				stack = stack[:len(stack)-k]
			}
		}
	}
	return string(stack)
}
