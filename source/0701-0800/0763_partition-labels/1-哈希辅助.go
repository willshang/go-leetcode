package main

import "fmt"

func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
}

// leetcode763_划分字母区间
func partitionLabels(S string) []int {
	m := make(map[byte]int)
	for i := 0; i < len(S); i++ {
		m[S[i]] = i
	}
	res := make([]int, 0)
	left := 0
	right := 0
	for i := 0; i < len(S); i++ {
		right = max(right, m[S[i]])
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
