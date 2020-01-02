package main

import (
	"fmt"
)

func main() {
	nums := []int{0, 3, 2, 2, 1, 1, 1, 2, 2, 1, 1}
	fmt.Println(majorityElement(nums))
}

func majorityElement(nums []int) int {
	halfCount := len(nums) / 2

	for _, v := range nums {
		countNum := 0
		for _, value := range nums {
			if v == value {
				countNum++
			}
		}
		if countNum > halfCount {
			return v
		}
	}
	return -1
}
