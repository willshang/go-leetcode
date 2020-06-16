package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(checkPossibility([]int{4,2,3}))
	//fmt.Println(checkPossibility([]int{4,2,1}))
	fmt.Println(checkPossibility([]int{1, 2, 3, 5, 4, 6}))
}

// leetcode665_非递减数列
func checkPossibility(nums []int) bool {
	count := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			if count == 1 {
				return false
			} else if i == 0 {
				// 4 2 3 => 2 2 3
				nums[i] = nums[i+1]
				count++
			} else if nums[i-1] > nums[i+1] {
				// 3 4 2 => 3 4 4
				nums[i+1] = nums[i]
				count++
			} else {
				// 1 4 2 =>  1 2 2
				nums[i] = nums[i+1]
				count++
			}
		}
	}
	return true
}
