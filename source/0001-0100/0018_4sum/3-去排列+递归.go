package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	//fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	//fmt.Println(fourSum([]int{0, 0}, 0))
	//fmt.Println(fourSum([]int{0, 1, 5, 0, 1, 5, 5, -4}, 11))
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}

// leetcode18_四数之和
var res [][]int

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res = make([][]int, 0)
	dfs(nums, target, []int{}, 0)
	return res
}

func dfs(nums []int, target int, arr []int, level int) {
	if len(arr) == 4 {
		sum := 0
		for i := 0; i < len(arr); i++ {
			sum = sum + arr[i]
		}
		if sum == target {
			tempArr := make([]int, len(arr))
			copy(tempArr, arr)
			res = append(res, tempArr)
		}
		return
	}
	prev := math.MaxInt32
	for i := level; i < len(nums); i++ {
		if nums[i] != prev {
			prev = nums[i]
			arr = append(arr, nums[i])
			dfs(nums, target, arr, i+1)
			arr = arr[:len(arr)-1]
		}
	}
}
