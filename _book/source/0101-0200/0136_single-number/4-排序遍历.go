package main

import (
	"fmt"
	"sort"
)

func main() {
	//arr := []int{4, 1, 2, 1, 2}
	arr := []int{2, 2, 1}
	fmt.Println(singleNumber(arr))
}

func singleNumber(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums); i = i + 2 {
		if i+1 == len(nums) {
			return nums[i]
		}
		if nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	return -1
}
