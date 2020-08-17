package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(increasingTriplet([]int{10, 7, 11, 8, 9, 10}))
	fmt.Println(increasingTriplet([]int{3, 4, 2, 6}))
}

func increasingTriplet(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] >= nums[j] {
				continue
			}
			for k := j + 1; k < len(nums); k++ {
				if nums[j] < nums[k] {
					return true
				}
			}
		}
	}
	return false
}
