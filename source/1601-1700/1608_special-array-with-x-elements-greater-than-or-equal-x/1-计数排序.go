package main

import "fmt"

func main() {
	fmt.Println(specialArray([]int{3, 6, 7, 7, 0}))
}

func specialArray(nums []int) int {
	arr := make([]int, 1001)
	for i := 0; i < len(nums); i++ {
		arr[nums[i]]++
	}
	for i := len(arr) - 2; i >= 0; i-- {
		arr[i] = arr[i] + arr[i+1]
	}
	for i := 0; i <= len(nums); i++ {
		if arr[i] == i {
			return i
		}
	}
	return -1
}
