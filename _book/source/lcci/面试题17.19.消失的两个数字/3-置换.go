package main

import "fmt"

func main() {
	fmt.Println(missingTwo([]int{2, 3}))
}

// 程序员面试金典17.19_消失的两个数字
func missingTwo(nums []int) []int {
	res := make([]int, 0)
	nums = append(nums, -1, -1, 0)
	for i := 0; i < len(nums); i++ {
		for nums[i] != -1 && nums[i] != i {
			nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
		}
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] == -1 {
			res = append(res, i)
		}
	}
	return res
}
