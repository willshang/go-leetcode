package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isValid("aabcbc"))
}

func isValid(s string) bool {
	for strings.Contains(s, "abc") {
		s = strings.ReplaceAll(s, "abc", "")
	}
	return s == ""
}
