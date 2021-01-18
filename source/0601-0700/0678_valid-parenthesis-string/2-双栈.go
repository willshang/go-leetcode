package main

import (
	"fmt"
)

func main() {
	fmt.Println(checkValidString("()"))
}

// leetcode678_有效的括号字符串
func checkValidString(s string) bool {
	stackL := make([]int, 0)
	stackS := make([]int, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stackL = append(stackL, i)
		} else if s[i] == '*' {
			stackS = append(stackS, i)
		} else {
			if len(stackL) > 0 {
				stackL = stackL[:len(stackL)-1]
			} else if len(stackS) > 0 {
				stackS = stackS[:len(stackS)-1]
			} else {
				return false
			}
		}
	}
	if len(stackL) > len(stackS) {
		return false
	}
	for len(stackL) > 0 && len(stackS) > 0 {
		a, b := stackL[len(stackL)-1], stackS[len(stackS)-1]
		if a > b {
			return false
		}
		stackL = stackL[:len(stackL)-1]
		stackS = stackS[:len(stackS)-1]
	}
	if len(stackL) == 0 {
		return true
	}
	return false
}
