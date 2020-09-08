package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumGap([]int{3, 6, 9, 1}))
}

// leetcode164_最大间距
func maximumGap(nums []int) int {
	res := 0
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] > res {
			res = nums[i] - nums[i-1]
		}
	}
	return res
}
