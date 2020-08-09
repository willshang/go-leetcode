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
	dfs(nums, "")
	return res
}

func dfs(nums []byte, str string) {
	if len(nums) == 0 {
		res = append(res, str)
		return
	}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		str = str + string(nums[i])
		arr := append([]byte{}, nums[:i]...)
		arr = append(arr, nums[i+1:]...)
		dfs(arr, str)
		str = str[:len(str)-1]
	}
}
