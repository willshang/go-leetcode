package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{2, 6, 4, 8, 10, 9, 15}
	fmt.Println(findUnsortedSubarray(arr))
}

func findUnsortedSubarray(nums []int) int {
	temp := make([]int, len(nums))
	copy(temp, nums)
	sort.Ints(temp)
	i, j := 0, len(nums)-1
	for i < len(nums) && nums[i] == temp[i] {
		i++
	}
	for i+1 < j && nums[j] == temp[j] {
		j--
	}
	return j - i + 1
}
