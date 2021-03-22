package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumSize([]int{9}, 2))
}

// leetcode1760_袋子里最少数目的球
func minimumSize(nums []int, maxOperations int) int {
	left := 1
	right := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > right {
			right = nums[i]
		}
	}
	res := 0
	for left <= right {
		mid := (left + right) / 2
		count := getCount(nums, mid)
		if count <= maxOperations {
			res = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return res
}

func getCount(nums []int, per int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		count = count + (nums[i]-1)/per
	}
	return count
}
