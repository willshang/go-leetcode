package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}

func minSubArrayLen(s int, nums []int) int {
	res := math.MaxInt32
	arr := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		arr[i] = arr[i-1] + nums[i-1]
	}
	for i := 1; i <= len(nums); i++ {
		target := s + arr[i-1]
		index := sort.SearchInts(arr, target)
		if index <= len(nums) {
			if res > index-i+1 {
				res = index - i + 1
			}
		}
	}
	if res == math.MaxInt32 {
		return 0
	}
	return res
}
