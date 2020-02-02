package main

import "fmt"

func main() {
	fmt.Println(isValid("()[]{}"))
}

// leetcode 20 有效的括号
func isValid(s string) bool {
	if s == "" {
		return true
	}

	stack := make([]rune, len(s))
	length := 0
	var match = map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		switch char {
		case '(', '[', '{':
			stack[length] = char
			length++
		case ')', ']', '}':
			if length == 0 {
				return false
			}
			if stack[length-1] != match[char] {
				return false
			} else {
				length--
			}
		}
	}
	return length == 0
}
