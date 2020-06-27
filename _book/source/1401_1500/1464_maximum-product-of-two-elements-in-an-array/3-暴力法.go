package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxProduct([]int{3, 4, 5, 2}))
}

func maxProduct(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if (nums[i]-1)*(nums[j]-1) > res {
				res = (nums[i] - 1) * (nums[j] - 1)
			}
		}
	}
	return res
}
