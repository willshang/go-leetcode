package main

import "fmt"

func main() {
	arr := []int{1, 7, 3, 6, 5, 6}
	fmt.Println(pivotIndex(arr))
}

func pivotIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	arr := make([]int, len(nums))
	arr[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		arr[i] = arr[i-1] + nums[i]
	}
	for i := 0; i < len(nums); i++ {
		var left, right int
		if i == 0 {
			left = 0
		} else {
			left = arr[i-1]
		}
		r := i + 1
		if r > len(nums)-1 {
			right = 0
		} else {
			right = arr[len(nums)-1] - arr[i]
		}
		if left == right {
			return i
		}
	}
	return -1
}
