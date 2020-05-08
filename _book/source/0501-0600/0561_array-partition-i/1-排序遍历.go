package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(arrayPairSum([]int{1, 4, 3, 2}))
}

// leetcode561_数组拆分 I
func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	sum := 0
	for k, v := range nums {
		if k%2 == 0 {
			sum = sum + v
		}
	}
	return sum
}
