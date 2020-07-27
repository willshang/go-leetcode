package main

import (
	"fmt"
)

func main() {
	fmt.Println(missingNumber([]int{6, 0, 4, 2, 3, 5, 7, 1, 9}))
	//fmt.Println(missingNumber([]int{3, 0, 1}))
}

func missingNumber(nums []int) int {
	n := len(nums)
	// 假设index=n
	index := n
	for i := 0; i < n; {
		// nums[i]到指定位置i后往后走
		if i == nums[i] {
			i++
			continue
		}
		if nums[i] == n {
			index = i
			i++
			continue
		}
		nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
	}
	return index
}
