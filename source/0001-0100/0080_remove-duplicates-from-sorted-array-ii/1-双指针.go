package main

import "fmt"

func main() {
	arr := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	removeDuplicates(arr)
	fmt.Println(arr)
}

// leetcode80_删除排序数组中的重复项II
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	n := 2
	i := n
	for j := n; j < len(nums); j++ {
		if nums[i-n] != nums[j] {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}
