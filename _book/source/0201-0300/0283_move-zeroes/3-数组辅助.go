package main

import "fmt"

func main() {
	nums := []int{6, 0, 1, 0, 3, 12, 0, 7}
	moveZeroes(nums)
	fmt.Println(nums)
}

func moveZeroes(nums []int) {
	arr := make([]int, len(nums))
	count := 0
	for i := range nums {
		if nums[i] != 0 {
			arr[count] = nums[i]
			count++
		}
	}

	copy(nums, arr)
}
