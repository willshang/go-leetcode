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

// leetcode771_宝石与石头
func numJewelsInStones(J string, S string) int {
	res := 0
	for _, v := range J {
		res = res + strings.Count(S, string(v))
	}
	return res
}
