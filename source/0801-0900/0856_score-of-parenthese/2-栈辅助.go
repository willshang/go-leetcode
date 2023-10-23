package main

import "fmt"

func main() {
	fmt.Println(scoreOfParentheses("(()(()))"))
}

func scoreOfParentheses(S string) int {
	stack := make([]int, 0)
	stack = append(stack, 0)
	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			stack = append(stack, 0)
		} else {
			a, b := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, b+max(2*a, 1))
		}
	}
	return stack[0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
