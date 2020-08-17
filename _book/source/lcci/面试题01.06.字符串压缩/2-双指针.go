package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(compressString("aabcccccaaa"))
}

func compressString(S string) string {
	if len(S) <= 1 {
		return S
	}
	i := 0
	j := 0
	res := ""
	for j = 1; j < len(S); j++ {
		if S[i] != S[j] {
			res = res + string(S[i]) + strconv.Itoa(j-i)
			i = j
		}
	}
	res = res + string(S[i]) + strconv.Itoa(j-i)
	if len(res) >= len(S) {
		return S
	}
	return res
}
