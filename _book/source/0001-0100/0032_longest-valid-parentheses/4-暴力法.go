package main

import "fmt"

func main() {
	fmt.Println(longestValidParentheses("((())"))
}

func longestValidParentheses(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		count := 0
		for j := i; j < len(s); j++ {
			if s[j] == '(' {
				count++
			} else {
				count--
			}
			if count < 0 {
				break
			}
			if count == 0 {
				res = max(res, j+1-i)
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
