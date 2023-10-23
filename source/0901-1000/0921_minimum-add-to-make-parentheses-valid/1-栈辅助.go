package main

import "fmt"

func main() {
	fmt.Println(minAddToMakeValid("())("))
}

// leetcode921_使括号有效的最少添加
func minAddToMakeValid(S string) int {
	stack := make([]byte, 0)
	res := 0
	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			stack = append(stack, '(')
		} else {
			if len(stack) == 0 {
				res++
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return res + len(stack)
}
