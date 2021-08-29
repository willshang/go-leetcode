package main

import "sort"

func main() {

}

func rearrangeArray(nums []int) []int {
	n := len(nums)
	res := make([]int, 0)
	sort.Ints(nums)
	m := (n + 1) / 2
	// 小大小大小...
	// 后一半数：前后都小于当前
	// 前一半数：前后都大于当前
	for i := 0; i < m; i++ {
		res = append(res, nums[i])
		if i+m < n {
			res = append(res, nums[i+m])
		}
	}
	return res
}
