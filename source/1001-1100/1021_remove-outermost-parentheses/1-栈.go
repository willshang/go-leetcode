package main

import "fmt"

func main() {
	fmt.Println(removeOuterParentheses("(()())(())"))
}

func removeOuterParentheses(S string) string {
	if len(S) == 0 {
		return ""
	}
	res := ""
	stack := make([]byte, 0)
	stack = append(stack, S[0])
	last := 0
	for i := 1; i < len(S); i++ {
		if len(stack) > 0 && S[i] == ')' && stack[len(stack)-1] == '(' {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				res = res + S[last+1:i]
				last = i + 1
			}
		} else {
			stack = append(stack, S[i])
		}
	}
	return res
}
