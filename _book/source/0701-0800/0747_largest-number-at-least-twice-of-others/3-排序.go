package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 6, 1, 0}
	fmt.Println(dominantIndex(nums))
	fmt.Println(dominantIndex([]int{1, 0}))
}

func dominantIndex(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	temp := make([]int, len(nums))
	copy(temp, nums)
	sort.Ints(temp)
	maxValue := temp[len(temp)-1]
	if maxValue < 2*temp[len(temp)-2] {
		return -1
	}
	for i := 0; i < n; i++ {
		if nums[i] == maxValue {
			return i
		}
	}
	return -1
}
