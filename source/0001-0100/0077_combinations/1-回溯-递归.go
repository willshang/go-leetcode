package main

import "fmt"

func main() {
	fmt.Println(combine(4, 2))
}

var res [][]int

func combine(n int, k int) [][]int {
	res = make([][]int, 0)
	nums := make([]int, 0)
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}
	dfs(nums, 0, k)
	return res
}

func dfs(nums []int, index, k int) {
	if index == k {
		temp := make([]int, k)
		copy(temp, nums[:k])
		res = append(res, temp)
		return
	}
	for i := index; i < len(nums); i++ {
		if index == 0 || nums[i] > nums[index-1] {
			nums[i], nums[index] = nums[index], nums[i]
			dfs(nums, index+1, k)
			nums[i], nums[index] = nums[index], nums[i]
		}
	}
}
