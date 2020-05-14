package main

import (
	"fmt"
)

func main() {
	S := "a1b2"
	fmt.Println(letterCasePermutation(S))
}

func letterCasePermutation(S string) []string {
	size := len(S)
	if size == 0 {
		return []string{""}
	}

	lastByte := S[size-1]
	postfixs := make([]string, 1, 2)
	postfixs[0] = string(lastByte)

	if b, ok := check(lastByte); ok {
		postfixs = append(postfixs, string(b))
	}
	prefixs := letterCasePermutation(S[:size-1])

	res := make([]string, 0, len(prefixs)*len(prefixs))
	for _, pre := range prefixs {
		for _, post := range postfixs {
			res = append(res, pre+post)
		}
	}
	return res
}

func check(b byte) (byte, bool) {
	if 'a' <= b && b <= 'z' {
		return b + 'A' - 'a', true
	}
	if 'A' <= b && b <= 'Z' {
		return b + 'a' - 'A', true
	}
	return 0, false
}
