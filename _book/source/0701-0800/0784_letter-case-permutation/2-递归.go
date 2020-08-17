package main

import (
	"fmt"
)

func main() {
	fmt.Println(letterCasePermutation("a1b2"))
}

func letterCasePermutation(S string) []string {
	size := len(S)
	if size == 0 {
		return []string{""}
	}
	postfixs := make([]string, 1)
	lastByte := S[size-1]
	postfixs[0] = string(lastByte)
	if b, ok := check(lastByte); ok {
		postfixs = append(postfixs, string(b))
	}
	prefixs := letterCasePermutation(S[:size-1])
	res := make([]string, 0)
	for _, pre := range prefixs {
		for _, post := range postfixs {
			res = append(res, pre+post)
		}
	}
	return res
}

func check(b byte) (byte, bool) {
	if 'a' <= b && b <= 'z' {
		return b - 'a' + 'A', true
	}
	if 'A' <= b && b <= 'Z' {
		return b - 'A' + 'a', true
	}
	return 0, false
}
