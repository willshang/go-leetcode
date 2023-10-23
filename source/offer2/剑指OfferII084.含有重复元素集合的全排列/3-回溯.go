package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
	fmt.Println(permuteUnique([]int{1, 2, 3}))
	fmt.Println(permuteUnique([]int{0, 1, 0, 0, 9}))
}

// 剑指OfferII084.含有重复元素集合的全排列
var res [][]int

func permuteUnique(nums []int) [][]int {
	res = make([][]int, 0)
	sort.Ints(nums)
	dfs(nums, make([]int, 0))
	return res
}

func dfs(nums []int, arr []int) {
	if len(nums) == 0 {
		temp := make([]int, len(arr))
		copy(temp, arr)
		res = append(res, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		tempArr := make([]int, len(nums))
		copy(tempArr, nums)
		arr = append(arr, nums[i])
		dfs(append(tempArr[:i], tempArr[i+1:]...), arr)
		arr = arr[:len(arr)-1]
	}
}
