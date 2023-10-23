package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(permutation("qqe"))
}

var res []string

func permutation(S string) []string {
	res = make([]string, 0)
	nums := []byte(S)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	dfs(nums, 0, make([]int, len(nums)), "")
	return res
}

func dfs(nums []byte, index int, visited []int, str string) {
	if len(nums) == index {
		res = append(res, str)
		return
	}
	for i := 0; i < len(nums); i++ {
		if visited[i] == 1 {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] && visited[i-1] == 0 {
			continue
		}
		str = str + string(nums[i])
		visited[i] = 1
		dfs(nums, index+1, visited, str)
		visited[i] = 0
		str = str[:len(str)-1]
	}
}
