package main

import "fmt"

func main() {
	fmt.Println(maxDepth("(1)+((2))+(((3)))"))
	fmt.Println(maxDepth("(1+(2*3)+((8)/4))+1"))
}

// leetcode1614_括号的最大嵌套深度
func maxDepth(s string) int {
	res := 0
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, s[i])
		} else if s[i] == ')' {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > res {
			res = len(stack)
		}
	}
	return res
}
