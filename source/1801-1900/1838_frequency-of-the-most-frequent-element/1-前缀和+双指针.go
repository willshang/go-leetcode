package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxFrequency([]int{1, 2, 4}, 5))
}

// leetcode1838_最高频元素的频数
func maxFrequency(nums []int, k int) int {
	n := len(nums)
	sort.Ints(nums)
	arr := make([]int, n+1)
	for i := 1; i <= n; i++ {
		arr[i] = arr[i-1] + nums[i-1]
	}
	res := 1
	i := 0
	for j := 0; j < n; j++ {
		for nums[j]*(j-i)-(arr[j]-arr[i]) > k {
			i++
		}
		if j-i+1 > res {
			res = j - i + 1
		}
	}
	return res
}
