package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(singleNumber([]int{2, 2, 3, 2}))
}

func singleNumber(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i = i + 3 {
		if nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	return nums[len(nums)-1]
}
