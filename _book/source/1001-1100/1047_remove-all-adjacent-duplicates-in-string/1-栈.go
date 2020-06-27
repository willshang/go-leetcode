package main

import (
	"fmt"
)

func main() {
	fmt.Println(removeDuplicates("abbaca"))
}

// leetcode1047_删除字符串中的所有相邻重复项
func removeDuplicates(S string) string {
	stack := make([]int32, 0)
	for _, v := range S {
		stack = append(stack, v)
		for len(stack) > 1 && stack[len(stack)-1] == stack[len(stack)-2] {
			stack = stack[:len(stack)-2]
		}
	}
	return string(stack)
}
