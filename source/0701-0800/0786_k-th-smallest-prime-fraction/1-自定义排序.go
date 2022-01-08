package main

import "sort"

func main() {

}

// leetcode786_第K个最小的素数分数
func kthSmallestPrimeFraction(arr []int, k int) []int {
	n := len(arr)
	nums := make([][]int, 0)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			nums = append(nums, []int{arr[i], arr[j]})
		}
	}
	sort.Slice(nums, func(i, j int) bool {
		// a/b c/d => ad<bc
		a, b, c, d := nums[i][0], nums[i][1], nums[j][0], nums[j][1]
		return a*d < b*c
	})
	return nums[k-1]
}
