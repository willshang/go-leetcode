package main

import (
	"fmt"
)

func main() {
	//fmt.Println(findRepeatNumber([]int{2, 3, 1, 0, 2, 5, 3}))
	fmt.Println(findRepeatNumber([]int{0, 1, 0}))
}

func findRepeatNumber(nums []int) int {
	countZero := 0
	for _, value := range nums {
		if value == 0 {
			if countZero > 0 {
				return 0
			}
			countZero++
			continue
		}
		if value < 0 {
			value = -value
		}
		if nums[value] < 0 {
			return value
		}
		nums[value] = -1 * nums[value]
	}
	return -1
}
