package main

import "fmt"

func main() {
	arr := []int{1, 7, 3, 6, 5, 6}
	fmt.Println(pivotIndex(arr))
}
func pivotIndex(nums []int) int {
	sum := 0
	for i := range nums {
		sum = sum + nums[i]
	}

	left := 0
	for i := range nums {
		if left*2+nums[i] == sum {
			return i
		}
		left = left + nums[i]
	}
	return -1
}
