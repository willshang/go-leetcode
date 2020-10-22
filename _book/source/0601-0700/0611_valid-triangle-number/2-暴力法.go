package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(triangleNumber([]int{2, 2, 3, 4}))
}

func triangleNumber(nums []int) int {
	sort.Ints(nums)
	res := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if nums[i]+nums[j] > nums[k] {
					res++
				} else {
					break
				}
			}
		}
	}
	return res
}
