package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(rangeSum([]int{1, 2, 3, 4}, 4, 1, 5))
}

// leetcode1508_子数组和排序后的区间和
func rangeSum(nums []int, n int, left int, right int) int {
	arr := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum = sum + nums[j]
			arr = append(arr, sum)
		}
	}
	sort.Ints(arr)
	res := 0
	for i := left - 1; i <= right-1; i++ {
		res = (res + arr[i]) % 1000000007
	}
	return res
}
