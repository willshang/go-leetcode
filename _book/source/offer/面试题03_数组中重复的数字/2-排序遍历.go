package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	fmt.Println(findRepeatNumber([]int{2, 3, 1, 0, 2, 5, 3}))
}

func findRepeatNumber(nums []int) int {
	sort.Ints(nums)
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == prev {
			return nums[i]
		}
		prev = nums[i]
	}
	return -1
}
