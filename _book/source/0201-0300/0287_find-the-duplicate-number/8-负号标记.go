package main

import (
	"fmt"
)

func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 4}))
}

func findDuplicate(nums []int) int {
	for i := 0; i < len(nums); i++ {
		index := abs(nums[i]) - 1
		if nums[index] > 0 {
			nums[index] = -1 * nums[index]
		} else {
			return abs(nums[i])
		}
	}
	return 0
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
