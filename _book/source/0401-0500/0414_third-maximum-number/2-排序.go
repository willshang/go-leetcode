package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{2, 2, 3, 1, 1}
	fmt.Println(thirdMax(nums))
}

func thirdMax(nums []int) int {

	sort.Ints(nums)
	if len(nums) < 3 {
		return nums[len(nums)-1]
	}

	k := 2
	maxValue := nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] != nums[i+1] {
			k--
		}
		if k == 0 {
			return nums[i]
		}
	}
	return maxValue
}
