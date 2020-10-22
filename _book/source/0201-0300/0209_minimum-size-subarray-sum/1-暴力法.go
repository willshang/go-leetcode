package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}

func minSubArrayLen(s int, nums []int) int {
	res := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum = sum + nums[j]
			if sum >= s {
				if res > j-i+1 {
					res = j - i + 1
				}
				break
			}
		}
	}
	if res == math.MaxInt32 {
		return 0
	}
	return res
}
