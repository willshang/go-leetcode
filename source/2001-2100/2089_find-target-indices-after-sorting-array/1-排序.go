package main

import "sort"

func main() {

}

// leetcode2089_找出数组排序后的目标下标
func targetIndices(nums []int, target int) []int {
	sort.Ints(nums)
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			res = append(res, i)
		}
	}
	return res
}
