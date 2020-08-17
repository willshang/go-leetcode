package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(missingNumber([]int{0, 1, 3}))
}

func missingNumber(nums []int) int {
	return sort.Search(len(nums), func(i int) bool {
		return nums[i] != i
	})
}
