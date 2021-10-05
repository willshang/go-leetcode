package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

// 剑指OfferII083.没有重复元素集合的全排列
var res [][]int

func permute(nums []int) [][]int {
	res = make([][]int, 0)
	arr := make([]int, 0)
	visited := make(map[int]bool)
	dfs(nums, 0, arr, visited)
	return res
}

func dfs(nums []int, index int, arr []int, visited map[int]bool) {
	if index == len(nums) {
		temp := make([]int, len(arr))
		copy(temp, arr)
		res = append(res, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if visited[i] == false {
			arr = append(arr, nums[i])
			visited[i] = true
			dfs(nums, index+1, arr, visited)
			arr = arr[:len(arr)-1]
			visited[i] = false
		}
	}
}
