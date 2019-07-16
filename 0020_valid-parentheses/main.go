package main

import (
	"fmt"
)

func main() {
	fmt.Println(isValid("()[]{}"))
}

func isValid(s string) bool {
	st := new(stack)
	for _, char := range s {
		switch char {
		case '(', '[', '{':
			st.push(char)
		case ')', ']', '}':
			ret, ok := st.pop()
			if !ok || ret != match[char] {
				return false
			}
		}
	}

	if len(*st) > 0 {
		return false
	}
	return true
}

var match = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
}

type stack []rune

func (s *stack) push(b rune) {
	*s = append(*s, b)
}
func (s *stack) pop() (rune, bool) {
	if len(*s) > 0 {
		res := (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		return res, true
	}
	return 0, false
}
