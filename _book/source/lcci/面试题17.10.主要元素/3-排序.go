package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(majorityElement([]int{3, 2}))
}

func majorityElement(nums []int) int {
	sort.Ints(nums)
	for i := 0; i <= len(nums)/2; i++ {
		if nums[i] == nums[i+len(nums)/2] {
			return nums[i]
		}
	}
	return -1
}
