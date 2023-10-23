package main

import (
	"fmt"
	"strings"
)

func main() {
	J := "aA"
	S := "aAAbbbb"
	fmt.Println(numJewelsInStones(J, S))
}

func numJewelsInStones(J string, S string) int {
	res := 0
	for _, v := range J {
		for _, s := range S {
			if v == s {
				res++
			}
		}
	}
	return res
}
