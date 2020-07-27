package main

import "fmt"

func main() {
	fmt.Println(balancedStringSplit("RLRRLLRLRL"))
}

func balancedStringSplit(s string) int {
	res := 0
	if len(s) == 0 {
		return res
	}
	stack := make([]byte, 0)
	stack = append(stack, s[0])
	for i := 1; i < len(s); i++ {
		if len(stack) > 0 &&
			((s[i] == 'L' && stack[len(stack)-1] == 'R') ||
				(s[i] == 'R' && stack[len(stack)-1] == 'L')) {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				res++
			}
		} else {
			stack = append(stack, s[i])
		}
	}
	return res
}
