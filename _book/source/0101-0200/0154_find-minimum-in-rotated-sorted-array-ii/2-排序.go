package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findMin([]int{2, 2, 2, 0, 1}))
}

func findMin(nums []int) int {
	sort.Ints(nums)
	return nums[0]
}
