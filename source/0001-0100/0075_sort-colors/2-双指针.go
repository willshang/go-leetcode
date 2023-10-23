package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 0, 2, 1, 1, 0}
	sortColors(arr)
	fmt.Println(arr)

}

// leetcode75_颜色分类
func sortColors(nums []int) {
	left := 0
	right := len(nums) - 1
	for i := 0; i <= right; i++ {
		if nums[i] == 0 {
			nums[left], nums[i] = nums[i], nums[left]
			left++
		} else if nums[i] == 2 {
			nums[right], nums[i] = nums[i], nums[right]
			right--
			i--
		}
	}
}
