package main

import "fmt"

func main() {
	fmt.Println(removeInvalidParentheses("()())()"))
}

func removeInvalidParentheses(s string) []string {

}

func isValid(s string) bool {
	left := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else if s[i] == ')' {
			left--
		}
		if left < 0 {
			return false
		}
	}
	return left == 0
}
