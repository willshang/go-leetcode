package main

import (
	"fmt"
)

func main() {
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
	fmt.Println(missingNumber([]int{3, 0, 1}))
}

func missingNumber(nums []int) int {
	n := len(nums)
	index := n
	for i := 0; i < n; {
		if nums[i] == n {
			index = i
			i++
			continue
		}
		if i == nums[i] {
			i++
			continue
		}
		nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
	}
	return index
}
