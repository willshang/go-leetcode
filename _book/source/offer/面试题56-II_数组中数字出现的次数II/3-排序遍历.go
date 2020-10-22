package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(singleNumber([]int{3, 4, 3, 3}))
}

func singleNumber(nums []int) int {
	sort.Ints(nums)
	if len(nums) == 1 {
		return nums[0]
	}
	i := 0
	for i < len(nums) {
		if i == len(nums)-1 {
			return nums[len(nums)-1]
		}
		if nums[i] == nums[i+1] {
			i = i + 3
		} else {
			return nums[i]
		}
	}
	return nums[i]
}
