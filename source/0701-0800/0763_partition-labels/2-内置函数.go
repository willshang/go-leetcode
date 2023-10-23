package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
}

func partitionLabels(S string) []int {
	res := make([]int, 0)
	left := 0
	right := 0
	for i := 0; i < len(S); i++ {
		right = max(right, strings.LastIndex(S, string(S[i])))
		if i == right {
			res = append(res, right-left+1)
			left = right + 1
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
