package main

import (
	"fmt"
	"strings"
)

func main() {
	//fmt.Println(letterCasePermutation("a1b2"))
	fmt.Println(letterCasePermutation("abcd"))
}

func letterCasePermutation(S string) []string {
	S = strings.ToLower(S)
	res := []string{S}
	for i := range S {
		if S[i] >= 'a' {
			n := len(res)
			for j := 0; j < n; j++ {
				temp := []byte(res[j])
				temp[i] = S[i] - 'a' + 'A'
				res = append(res, string(temp))
			}
		}
	}
	return res
}
