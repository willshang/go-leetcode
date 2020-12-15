package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(minRemoveToMakeValid("lee(t(c)o)de)"))
}

// leetcode1249_移除无效的括号
func minRemoveToMakeValid(s string) string {
	arr := []byte(s)
	stack := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		if arr[i] == '(' {
			stack = append(stack, i)
		} else if arr[i] == ')' {
			if len(stack) == 0 {
				arr[i] = ' '
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	for i := 0; i < len(stack); i++ {
		arr[stack[i]] = ' '
	}
	return strings.ReplaceAll(string(arr), " ", "")
}
