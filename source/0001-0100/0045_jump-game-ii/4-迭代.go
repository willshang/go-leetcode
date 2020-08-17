package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
	//fmt.Println(jump([]int{2, 1}))
}

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	arr := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		arr[i] = math.MaxInt32
	}
	arr[0] = 0
	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] >= len(nums)-1 {
			return arr[i] + 1
		}
		for j := i + 1; j <= i+nums[i]; j++ {
			if j < len(nums) {
				arr[j] = min(arr[j], arr[i]+1)
			} else {
				break
			}
		}
	}
	return arr[len(nums)-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
