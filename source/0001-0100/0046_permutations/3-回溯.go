package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

var res [][]int

func permute(nums []int) [][]int {
	res = make([][]int, 0)
	arr := make([]int, len(nums))
	dfs(nums, 0, arr)
	return res
}

func dfs(nums []int, index int, arr []int) {
	if index == len(nums) {
		temp := make([]int, len(arr))
		copy(temp, arr)
		res = append(res, temp)
		return
	}
	for i := index; i < len(nums); i++ {
		arr[index] = nums[i]
		nums[i], nums[index] = nums[index], nums[i]
		dfs(nums, index+1, arr)
		nums[i], nums[index] = nums[index], nums[i]
	}
}
