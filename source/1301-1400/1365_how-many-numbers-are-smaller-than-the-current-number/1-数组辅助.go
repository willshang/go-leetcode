package main

import "fmt"

func main() {
	fmt.Println(smallerNumbersThanCurrent([]int{8, 1, 2, 2, 3}))
}

func smallerNumbersThanCurrent(nums []int) []int {
	arr := make([]int, 101)
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		arr[nums[i]]++
	}
	for i := 1; i < len(arr); i++ {
		arr[i] = arr[i] + arr[i-1]
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			res[i] = arr[nums[i]-1]
		}
	}
	return res
}
