package main

import "fmt"

func main() {
	fmt.Println(removeOuterParentheses("(()())(())"))
}

// leetcode1021_删除最外层的括号
func removeOuterParentheses(S string) string {
	res := ""
	stack := make([]byte, 0)
	for i := 0; i < len(S); i++ {
		if S[i] == ')' {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			res = res + string(S[i])
		}
		if S[i] == '(' {
			stack = append(stack, S[i])
		}
	}
	return res
}
