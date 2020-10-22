package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(minAddToMakeValid("())("))
}

func minAddToMakeValid(S string) int {
	for strings.Contains(S, "()") == true {
		S = strings.ReplaceAll(S, "()", "")
	}
	return len(S)
}
