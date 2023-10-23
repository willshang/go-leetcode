package main

import (
	"fmt"
)

func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 4}))
}

// leetcode287_寻找重复数
func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]
	for slow != fast {
		slow, fast = nums[slow], nums[nums[fast]]
	}
	slow = 0
	for slow != fast {
		slow, fast = nums[slow], nums[fast]
	}
	return slow
}
