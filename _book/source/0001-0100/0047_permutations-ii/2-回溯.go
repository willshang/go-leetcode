package main

import "fmt"

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
}

var res [][]int

func permuteUnique(nums []int) [][]int {
	res = make([][]int, 0)
	dfs(nums, 0)
	return res
}

func dfs(nums []int, index int) {
	if index == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, nums)
		res = append(res, temp)
		return
	}
	for i := 0; i < index; i++ {
		if nums[i] == nums[index] {
			return
		}
		nums[i], nums[index] = nums[index], nums[i]
		dfs(nums, index+1)
		nums[i], nums[index] = nums[index], nums[i]
	}
}
