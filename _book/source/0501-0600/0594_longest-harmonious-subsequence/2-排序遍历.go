package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 3, 2, 2, 5, 2, 3, 7}
	fmt.Println(findLHS(arr))
	fmt.Println(findLHS([]int{1, 1, 1, 1}))
}

func findLHS(nums []int) int {
	sort.Ints(nums)
	res := 0
	left := 0
	for i := 0; i < len(nums); i++ {
		for nums[i]-nums[left] > 1 {
			left++
		}
		if nums[i]-nums[left] == 1 {
			if res < i-left+1 {
				res = i - left + 1
			}
		}
	}
	return res
}
