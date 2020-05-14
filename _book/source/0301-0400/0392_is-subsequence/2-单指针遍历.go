package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc"))
}

func isSubsequence(s string, t string) bool {
	for _, v := range s {
		idx := strings.IndexRune(t, v)
		if idx == -1 {
			return false
		}
		t = t[idx+1:]
	}
	return true
}
