package main

import (
	"fmt"
)

func main() {
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))
}

func firstMissingPositive(nums []int) int {
	n := len(nums)
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		m[nums[i]] = 1
	}
	for i := 1; i <= n; i++ {
		if m[i] == 0 {
			return i
		}
	}
	return n + 1
}
