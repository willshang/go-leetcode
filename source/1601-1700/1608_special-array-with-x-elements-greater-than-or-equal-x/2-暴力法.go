package main

import "fmt"

func main() {
	fmt.Println(specialArray([]int{3, 6, 7, 7, 0}))
}

// leetcode1608_特殊数组的特征值
func specialArray(nums []int) int {
	for i := 0; i <= len(nums); i++ {
		count := 0
		for j := 0; j < len(nums); j++ {
			if nums[j] >= i {
				count++
			}
		}
		if count == i {
			return i
		}
	}
	return -1
}
