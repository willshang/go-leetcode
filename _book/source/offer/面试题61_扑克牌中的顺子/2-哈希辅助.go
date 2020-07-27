package main

import (
	"fmt"
)

func main() {
	fmt.Println(isStraight([]int{1, 2, 3, 4, 5}))
	fmt.Println(isStraight([]int{0, 0, 1, 2, 5}))
}

func isStraight(nums []int) bool {
	m := make(map[int]bool)
	max, min := -1, 14
	countZero := 0
	for i := 0; i < 5; i++ {
		if nums[i] == 0 {
			countZero++
			continue
		}
		if m[nums[i]] {
			return false
		}
		m[nums[i]] = true
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}
	}
	if countZero == 0 {
		return max-min == 4
	}
	return max-min <= 4
}
