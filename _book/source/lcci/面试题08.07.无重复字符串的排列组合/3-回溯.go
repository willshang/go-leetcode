package main

import "fmt"

func main() {
	fmt.Println(permutation("qwe"))
}

// 程序员面试金典08.07_无重复字符串的排列组合
var res []string

func permutation(S string) []string {
	res = make([]string, 0)
	nums := []byte(S)
	dfs(nums, 0, "")
	return res
}

func dfs(nums []byte, index int, str string) {
	if index == len(nums) {
		res = append(res, str)
		return
	}
	for i := index; i < len(nums); i++ {
		str = str + string(nums[i])
		nums[i], nums[index] = nums[index], nums[i]
		dfs(nums, index+1, str)
		nums[i], nums[index] = nums[index], nums[i]
		str = str[:len(str)-1]
	}
}
