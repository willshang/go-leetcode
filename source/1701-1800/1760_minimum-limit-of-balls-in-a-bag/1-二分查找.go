package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumSize([]int{9}, 2))
}

func minimumSize(nums []int, maxOperations int) int {
	sort.Ints(nums)
	left := 1
	right := nums[len(nums)-1]
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
		if nums[i]%per == 0 {
			count = count + nums[i]/per - 1
		} else {
			count = count + nums[i]/per
		}
	}
	return count
}
