package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(numDifferentIntegers("0a0"))
}

func numDifferentIntegers(word string) int {
	m := make(map[string]bool)
	arr := strings.FieldsFunc(word, func(r rune) bool {
		return 'a' <= r && r <= 'z'
	})
	for i := 0; i < len(arr); i++ {
		s := strings.Trim(arr[i], " ")
		for len(s) > 0 && s[0] == '0' {
			s = s[1:]
		}
		m[s] = true
	}
	return len(m)
}
