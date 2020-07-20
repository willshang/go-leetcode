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
	dfs(nums, make([]int, k), 0, k)
	return res
}

func dfs(nums []int, arr []int, index, k int) {
	if index == k {
		temp := make([]int, len(arr))
		copy(temp, arr)
		res = append(res, temp)
		return
	}
	for i := index; i < len(nums); i++ {
		arr[index] = nums[i]
		nums[i], nums[index] = nums[index], nums[i]
		dfs(nums, arr, index+1, k)
		nums[i], nums[index] = nums[index], nums[i]
	}
}
