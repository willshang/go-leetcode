package main

import (
	"fmt"
)

func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 4}))
}

func findDuplicate(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	for i := 1; i <= len(nums)-1; i++ {
		for i != nums[i] {
			value := nums[i]
			if value == nums[value] {
				return value
			}
			nums[i], nums[value] = nums[value], nums[i]
		}
	}
	return nums[0]
}
