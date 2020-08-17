package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}

// leetcode209_长度最小的子数组
func minSubArrayLen(s int, nums []int) int {
	res := math.MaxInt32
	i, j := 0, 0
	sum := 0
	for ; j < len(nums); j++ {
		sum = sum + nums[j]
		for sum >= s {
			if res > j-i+1 {
				res = j - i + 1
			}
			sum = sum - nums[i]
			i++
		}
	}
	if res == math.MaxInt32 {
		return 0
	}
	return res
}
