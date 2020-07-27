package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(reverseOnlyLetters("ab-cd"))
}

func reverseOnlyLetters(S string) string {
	stack := make([]rune, 0)
	res := make([]rune, 0)
	arr := []rune(S)
	for i := 0; i < len(arr); i++ {
		if unicode.IsLetter(arr[i]) {
			stack = append(stack, arr[i])
		}
	}
	for i := 0; i < len(arr); i++ {
		if unicode.IsLetter(arr[i]) {
			res = append(res, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		} else {
			res = append(res, arr[i])
		}
	}
	return string(res)
}
