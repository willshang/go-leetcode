package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(checkPossibility([]int{4,2,3}))
	//fmt.Println(checkPossibility([]int{4,2,1}))
	fmt.Println(checkPossibility([]int{1, 2, 3, 5, 4, 6}))
}

// 3 4 2 => 3 4 4 | 3 2 2
// 1 4 2 => 1 4 4 | 1 2 2
// 4 2 3 => 2 2 3 | 4 4 3
func checkPossibility(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			pre := deepCopy(nums)
			pre[i-1] = pre[i]
			next := deepCopy(nums)
			next[i] = next[i-1]
			return sort.IsSorted(sort.IntSlice(pre)) || sort.IsSorted(sort.IntSlice(next))
		}
	}
	return true
}

func deepCopy(nums []int) []int {
	res := make([]int, len(nums))
	copy(res, nums)
	return res
}
