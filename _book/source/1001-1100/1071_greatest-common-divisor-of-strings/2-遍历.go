package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(gcdOfStrings("ABCABC", "ABC"))
}

func gcdOfStrings(str1 string, str2 string) string {
	min := len(str1)
	if min > len(str2) {
		min = len(str2)
	}
	for i := len(str2); i >= 1; i-- {
		if len(str1)%i == 0 && len(str2)%i == 0 && str1[:i] == str2[:i] {
			a := strings.Repeat(str1[:i], len(str1)/i)
			b := strings.Repeat(str2[:i], len(str2)/i)
			if a == str1 && b == str2 {
				return str1[:i]
			}
		}
	}
	return ""
}
