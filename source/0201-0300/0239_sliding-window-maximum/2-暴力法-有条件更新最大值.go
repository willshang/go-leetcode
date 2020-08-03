package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}

func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0)
	if len(nums) == 0 {
		return res
	}
	max := math.MaxInt32
	for i := 0; i < len(nums)-k+1; i++ {
		if i == 0 || nums[i-1] == max {
			max = nums[i]
			for j := i; j < i+k; j++ {
				if nums[j] > max {
					max = nums[j]
				}
			}
		} else {
			if nums[i+k-1] > max {
				max = nums[i+k-1]
			}
		}
		res = append(res, max)
	}
	return res
}
