package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(specialArray([]int{3, 6, 7, 7, 0}))
}

func specialArray(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	if nums[0] >= n {
		return n
	}
	for i := 1; i < n; i++ {
		target := n - i
		if nums[i] >= target && target > nums[i-1] {
			return target
		}
	}
	return -1
}
