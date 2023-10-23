package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxSumRangeQuery([]int{1, 2, 3, 4, 5}, [][]int{{1, 3}, {0, 1}}))
}

// leetcode1589_所有排列中的最大和
func maxSumRangeQuery(nums []int, requests [][]int) int {
	d := make([]int, len(nums)+1)
	arr := make([]int, len(nums))
	for i := 0; i < len(requests); i++ {
		start := requests[i][0]
		end := requests[i][1]
		d[start]++
		d[end+1]--
	}
	arr[0] = d[0]
	for i := 1; i < len(nums); i++ {
		arr[i] = d[i] + arr[i-1]
	}
	sort.Ints(arr)
	sort.Ints(nums)
	res := 0
	for i := len(arr) - 1; i >= 0; i-- {
		res = (res + arr[i]*nums[i]) % 1000000007
	}
	return res
}
