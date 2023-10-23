package main

import "sort"

func main() {

}

// leetcode1968_构造元素不等于两相邻元素平均值的数组
func rearrangeArray(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	sort.Ints(nums)
	index := 0
	// 小大小大小...
	// 后一半数：前后都小于当前
	// 前一半数：前后都大于当前
	for i := 0; i < n; i++ {
		res[index] = nums[i]
		index = index + 2
		if index >= n {
			index = 1
		}
	}
	return res
}
