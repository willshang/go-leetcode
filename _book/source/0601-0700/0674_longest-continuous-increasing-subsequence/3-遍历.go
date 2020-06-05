package main

import "fmt"

func main() {
	arr := []int{1, 3, 5, 4, 7}
	fmt.Println(findLengthOfLCIS(arr))
}

// leetcode674_最长连续递增序列
func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := 1
	count := 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			count++
		} else {
			count = 1
		}
		if count > res {
			res = count
		}
	}
	return res
}
