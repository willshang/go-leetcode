package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(reverseOnlyLetters("ab-cd"))
}

// leetcode917_仅仅反转字母
func reverseOnlyLetters(S string) string {
	i := 0
	j := len(S) - 1
	arr := []rune(S)
	for i < j {
		if !unicode.IsLetter(arr[i]) {
			i++
		} else if !unicode.IsLetter(arr[j]) {
			j--
		} else {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}
	return string(arr)
}
