package main

import "fmt"

func main() {
	arr := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	removeDuplicates(arr)
	fmt.Println(arr)
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	prev := nums[0]
	count := 1
	j := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == prev {
			count = count + 1
		} else {
			count = 1
			prev = nums[i]
		}
		if count <= 2 {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}
