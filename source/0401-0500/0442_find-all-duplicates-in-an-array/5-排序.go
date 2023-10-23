package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}

func findDuplicates(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	sort.Ints(nums)
	res := make([]int, 0)
	prev := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if prev == nums[i] {
			count++
		} else {
			if count == 2 {
				res = append(res, nums[i-1])
			}
			prev = nums[i]
			count = 1
		}
	}
	if count == 2 {
		res = append(res, nums[len(nums)-1])
	}
	return res
}
