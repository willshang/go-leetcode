package main

import "fmt"

func main() {
	arr := []int{1, 1, 0, 1, 1, 1}
	fmt.Println(findMaxConsecutiveOnes(arr))
}

// leetcode485_最大连续1的个数
func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	count := 0
	for _, v := range nums {
		if v == 1 {
			count++
		} else {
			if count > max {
				max = count
			}
			count = 0
		}
	}
	if count > max {
		max = count
	}
	return max
}
