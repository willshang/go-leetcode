package main

import "fmt"

func main() {
	fmt.Println(expectNumber([]int{1, 2, 3}))
}

// leetcode_lcp11_期望个数统计
func expectNumber(scores []int) int {
	m := make(map[int]bool)
	for i := 0; i < len(scores); i++ {
		m[scores[i]] = true
	}
	return len(m)
}
