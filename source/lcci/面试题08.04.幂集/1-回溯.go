package main

import "fmt"

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

// 程序员面试金典08.04_幂集
var res [][]int

func subsets(nums []int) [][]int {
	res = make([][]int, 0)
	dfs(nums, make([]int, 0), 0)
	return res
}

func dfs(nums []int, arr []int, level int) {
	temp := make([]int, len(arr))
	copy(temp, arr)
	res = append(res, temp)
	for i := level; i < len(nums); i++ {
		// dfs(nums, append(arr, nums[i]), i+1)
		arr = append(arr, nums[i])
		dfs(nums, arr, i+1)
		arr = arr[:len(arr)-1]
	}
}
