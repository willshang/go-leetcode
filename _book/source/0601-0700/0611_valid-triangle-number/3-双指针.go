package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(triangleNumber([]int{0, 1, 1, 1}))
}

func triangleNumber(nums []int) int {
	sort.Ints(nums)
	res := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			continue
		}
		left, right := i+1, i+2
		for left < n-1 && nums[left] != 0 {
			for right < n && nums[i]+nums[left] > nums[right] {
				right++
			}
			res = res + right - left - 1
			left++
		}
	}
	return res
}
