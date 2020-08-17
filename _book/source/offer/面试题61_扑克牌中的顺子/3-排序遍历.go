package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(isStraight([]int{1, 2, 3, 4, 5}))
}

// 剑指offer_面试题61_扑克牌中的顺子
func isStraight(nums []int) bool {
	sort.Ints(nums)
	zero := 0
	gap := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zero++
		}
	}
	small := zero
	big := small + 1
	for big < len(nums) {
		if nums[small] == nums[big] {
			return false
		}
		gap = gap + nums[big] - nums[small] - 1
		small++
		big++
	}
	if gap > zero {
		return false
	}
	return true
}
