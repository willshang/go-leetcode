package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(canReorderDoubled([]int{3, 1, 3, 6}))
	//fmt.Println(canReorderDoubled([]int{4, -2, 2, -4}))
	fmt.Println(canReorderDoubled([]int{-6, -3}))
}

// leetcode954_二倍数对数组
func canReorderDoubled(A []int) bool {
	m := make(map[int]int)
	for i := 0; i < len(A); i++ {
		m[A[i]]++
	}
	sort.Slice(A, func(i, j int) bool {
		return abs(A[i]) < abs(A[j])
	})
	for i := 0; i < len(A); i++ {
		if m[A[i]] == 0 {
			continue
		}
		if m[2*A[i]] == 0 {
			return false
		}
		m[A[i]]--
		m[2*A[i]]--
	}
	return true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
