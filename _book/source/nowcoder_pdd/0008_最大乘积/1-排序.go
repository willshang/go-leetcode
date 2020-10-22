package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m int
	for {
		a, _ := fmt.Scan(&n)
		if a == 0 {
			break
		}
		nums := make([]int, 0)
		for i := 0; i < n; i++ {
			_, _ = fmt.Scan(&m)
			nums = append(nums, m)
		}
		fmt.Println(maximumProduct(nums))
	}
}

func maximumProduct(nums []int) int {
	sort.Ints(nums)
	return max(nums[0]*nums[1]*nums[len(nums)-1],
		nums[len(nums)-3]*nums[len(nums)-2]*nums[len(nums)-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
