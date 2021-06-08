package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(reductionOperations([]int{1, 1, 2, 2, 3}))
}

// leetcode1887_使数组元素相等的减少操作次数
func reductionOperations(nums []int) int {
	sort.Ints(nums)
	res := 0
	count := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			count++
		}
		res = res + count
	}
	return res
}
