package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 4}))
}

func findDuplicate(nums []int) int {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return nums[i]
		}
	}
	return -1
}
