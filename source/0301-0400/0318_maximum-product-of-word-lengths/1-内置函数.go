package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(maxProduct([]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}))
}

func maxProduct(words []string) int {
	res := 0
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if strings.ContainsAny(words[i], words[j]) == false &&
				res < len(words[i])*len(words[j]) {
				res = len(words[i]) * len(words[j])
			}
		}
	}
	return res
}
