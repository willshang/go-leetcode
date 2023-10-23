package main

import "fmt"

func main() {
	fmt.Println(majorityElement([]int{3, 2, 1}))
}

func majorityElement(nums []int) int {
	result, count := 0, 0
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			result = nums[i]
			count++
		} else if result == nums[i] {
			count++
		} else {
			count--
		}
	}
	total := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == result {
			total++
		}
	}
	if total <= len(nums)/2 {
		return -1
	}
	return result
}
