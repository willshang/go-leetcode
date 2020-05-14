package main

import (
	"fmt"
)

func main() {
	J := "aA"
	S := "aAAbbbb"
	fmt.Println(numJewelsInStones(J, S))
}

func numJewelsInStones(J string, S string) int {
	m := make(map[byte]bool)
	for i := range J {
		m[J[i]] = true
	}
	res := 0
	for i := range S {
		if m[S[i]] {
			res++
		}
	}
	return res
}
