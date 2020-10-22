package main

import "fmt"

func main() {
	fmt.Println(longestValidParentheses("((())"))
}

// leetcode32_最长有效括号
func longestValidParentheses(s string) int {
	res := 0
	stack := make([]int, 0)
	stack = append(stack, -1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1] // 弹出栈顶元素表示匹配了当前右括号
			if len(stack) == 0 {         // 没有匹配到左括号，存入最后一个没有被匹配到的右括号下标
				stack = append(stack, i)
			} else {
				res = max(res, i-stack[len(stack)-1])
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
