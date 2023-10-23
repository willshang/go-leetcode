package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(reductionOperations([]int{1, 1, 2, 2, 3}))
}

func reductionOperations(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	res := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			res = res + i
		}
	}
	return res
}
