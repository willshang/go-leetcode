package main

import (
	"fmt"
)

func main() {
	arr := []int{5, 2, 3, 1}
	sortArray(arr)
	fmt.Println(arr)
}

func sortArray(nums []int) []int {
	n := len(nums)
	for gap := n / 2; gap > 0; gap = gap / 2 {
		for i := gap; i < n; i++ {
			j := i
			cur := nums[i]
			for j-gap >= 0 && cur < nums[j-gap] {
				nums[j] = nums[j-gap]
				j = j - gap
			}
			nums[j] = cur
		}
	}
	return nums
}
