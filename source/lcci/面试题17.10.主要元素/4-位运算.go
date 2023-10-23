package main

import (
	"fmt"
)

func main() {
	fmt.Println(majorityElement([]int{3, 2}))
}

func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	result := int32(0)
	mask := int32(1)
	for i := 0; i < 32; i++ {
		count := 0
		for j := 0; j < len(nums); j++ {
			if mask&int32(nums[j]) == mask {
				count++
			}
		}
		if count > len(nums)/2 {
			result = result | mask
		}
		mask = mask << 1
	}
	total := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == int(result) {
			total++
		}
	}
	if total <= len(nums)/2 {
		return -1
	}
	return int(result)
}
