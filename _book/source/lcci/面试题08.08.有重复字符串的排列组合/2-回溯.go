package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(permutation("qqe"))
}

// 程序员面试金典08.08_有重复字符串的排列组合
var res []string

func permutation(S string) []string {
	res = make([]string, 0)
	nums := []byte(S)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	dfs(nums, 0)
	return res
}

func dfs(nums []byte, index int) {
	if index == len(nums) {
		res = append(res, string(nums))
		return
	}
	m := make(map[byte]int)
	for i := index; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			continue
		}
		m[nums[i]] = 1
		nums[i], nums[index] = nums[index], nums[i]
		dfs(nums, index+1)
		nums[i], nums[index] = nums[index], nums[i]
	}
}
