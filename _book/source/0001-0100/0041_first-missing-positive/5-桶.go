package main

import (
	"fmt"
)

func main() {
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))
}

func firstMissingPositive(nums []int) int {
	n := len(nums)
	arr := make([]int, n+1)
	for i := 0; i < n; i++ {
		if 1 <= nums[i] && nums[i] <= n {
			arr[nums[i]] = 1
		}
	}
	for i := 1; i <= n; i++ {
		if arr[i] == 0 {
			return i
		}
	}
	return n + 1
}
