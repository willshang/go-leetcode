package main

import (
	"fmt"
)

func main() {
	S := "a1b2"
	fmt.Println(letterCasePermutation(S))
}

func letterCasePermutation(S string) []string {
	res := make([]string, 1)
	for i := 0; i < len(S); i++ {
		if string(S[i]) >= "0" && string(S[i]) <= "9" {
			newRes := make([]string, 0)
			for _, v := range res {
				newRes = append(newRes, v+string(S[i]))
			}
			res = newRes
		} else if b, ok := check(S[i]); ok {
			first := string(b[0])
			second := string(b[1])
			newRes := make([]string, 0)
			for _, v := range res {
				newRes = append(newRes, v+first)
				newRes = append(newRes, v+second)
			}
			res = newRes
		}
	}
	return res
}

func check(b byte) ([]byte, bool) {
	if 'a' <= b && b <= 'z' {
		return []byte{b - 'a' + 'A', b}, true
	}
	if 'A' <= b && b <= 'Z' {
		return []byte{b, b - 'A' + 'a'}, true
	}
	return []byte{b}, false
}
