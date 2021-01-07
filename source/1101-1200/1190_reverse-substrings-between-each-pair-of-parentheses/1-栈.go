package main

import "strings"

func main() {

}

// leetcode1190_反转每对括号间的子串
func reverseParentheses(s string) string {
	arr := []byte(s)
	stack := make([]int, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else if s[i] == ')' {
			reverse(arr, stack[len(stack)-1], i)
			stack = stack[:len(stack)-1]
		}
	}
	s = string(arr)
	s = strings.ReplaceAll(s, "(", "")
	s = strings.ReplaceAll(s, ")", "")
	return s
}

func reverse(arr []byte, i, j int) {
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}
