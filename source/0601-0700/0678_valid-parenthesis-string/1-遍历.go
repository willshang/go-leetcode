package main

import "fmt"

func main() {
	fmt.Println(checkValidString("()"))
}

func checkValidString(s string) bool {
	// 第1次把星号当左括号看
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == ')' {
			right++
		} else {
			left++
		}
		if right > left {
			return false
		}
	}
	// 第2次把星号当右括号看
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left > right {
			return false
		}
	}
	return true
}
