package main

import (
	"fmt"
)

func main() {
	fmt.Println(minInsertions("(())"))
}

// leetcode1541_平衡括号字符串的最少插入次数
func minInsertions(s string) int {
	res := 0
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, '(')
		} else if s[i] == ')' {
			if i+1 < len(s) && s[i+1] == ')' {
				if len(stack) == 0 {
					// '))'的情况，补'('
					res++
				} else {
					stack = stack[:len(stack)-1]
				}
				i++
			} else {
				if len(stack) == 0 {
					// ')('的情况，')'补'()'
					res = res + 2
				} else {
					// '()('的情况, '()'补')'
					res++
					stack = stack[:len(stack)-1]
				}
			}
		}
	}
	res = res + len(stack)*2
	return res
}
