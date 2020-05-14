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
		if string(S[i]) <= "9" && string(S[i]) >= "0" {
			newres := []string{}
			for _, v := range res {
				newres = append(newres, v+string(S[i]))
			}
			res = newres
		} else if b, ok := check(S[i]); ok {
			first := string(b[0])
			second := string(b[1])
			newres := make([]string, 0)
			for _, v := range res {
				newres = append(newres, v+first)
				newres = append(newres, v+second)
			}
			res = newres
		}
	}
	return res
}

func check(b byte) ([]byte, bool) {
	if 'a' <= b && b <= 'z' {
		return []byte{b + 'A' - 'a', b}, true
	}
	if 'A' <= b && b <= 'Z' {
		return []byte{b, b + 'a' - 'A'}, true
	}
	return []byte{b}, false
}
