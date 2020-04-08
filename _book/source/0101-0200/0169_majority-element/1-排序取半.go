package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{0, 3, 2, 2, 1, 1, 1, 2, 2, 1, 1}
	fmt.Println(majorityElement(nums))
}

func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}
