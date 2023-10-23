package main

import "fmt"

func main() {
	fmt.Println(isValid("aabcbc"))
}

// leetcode1003_检查替换后的词是否有效
func isValid(s string) bool {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		stack = append(stack, s[i])
		if len(stack) >= 3 && string(stack[len(stack)-3:]) == "abc" {
			stack = stack[:len(stack)-3]
		}
	}
	return len(stack) == 0
}
