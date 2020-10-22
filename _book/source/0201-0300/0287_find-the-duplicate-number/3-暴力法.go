package main

import (
	"fmt"
)

func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 4}))
}

func findDuplicate(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return nums[i]
			}
		}
	}
	return -1
}
