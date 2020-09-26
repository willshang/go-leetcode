package main

import (
	"fmt"
)

func main() {
	nums := []int{0, 3, 2, 2, 1, 1, 1, 2, 2, 1, 1}
	fmt.Println(majorityElement(nums))
}

func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	result := int32(0)
	// 64位有坑
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
	return int(result)
}
