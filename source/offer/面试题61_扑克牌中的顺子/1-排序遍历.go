package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(isStraight([]int{1, 2, 3, 4, 5}))
}

func isStraight(nums []int) bool {
	sort.Ints(nums)
	sum := 0
	for i := 0; i < 4; i++ {
		if nums[i] == 0 {
			continue
		}
		// 非王重复
		if nums[i] == nums[i+1] {
			return false
		}
		sum = sum + nums[i+1] - nums[i]
	}
	return sum <= 4
}
