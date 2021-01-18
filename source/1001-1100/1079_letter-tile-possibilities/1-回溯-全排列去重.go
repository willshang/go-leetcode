package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(numTilePossibilities("AAB"))
}

// leetcode1079_活字印刷
var res [][]byte

func numTilePossibilities(tiles string) int {
	res = make([][]byte, 0)
	arr := []byte(tiles)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	dfs(arr, 0, make([]int, len(arr)), make([]byte, 0))
	return len(res)
}

func dfs(nums []byte, index int, visited []int, arr []byte) {
	if len(arr) > 0 {
		temp := make([]byte, len(arr))
		copy(temp, arr)
		res = append(res, temp)
	}
	for i := 0; i < len(nums); i++ {
		if visited[i] == 1 {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] && visited[i-1] == 0 {
			continue
		}
		arr = append(arr, nums[i])
		visited[i] = 1
		dfs(nums, index+1, visited, arr)
		visited[i] = 0
		arr = arr[:len(arr)-1]
	}
}
