package main

import (
	"fmt"
)

func main() {
	fmt.Println(letterCasePermutation("a1b2"))
}

// leetcode784_字母大小写全排列
var res []string

func letterCasePermutation(S string) []string {
	res = make([]string, 0)
	dfs([]byte(S), 0)
	return res
}

func dfs(arr []byte, level int) {
	if level == len(arr) {
		res = append(res, string(arr))
		return
	}
	if arr[level] >= 'a' && arr[level] <= 'z' {
		dfs(arr, level+1)
		arr[level] = arr[level] - 'a' + 'A' // 大写
		dfs(arr, level+1)
		arr[level] = arr[level] - 'A' + 'a' // 小写
	} else if arr[level] >= 'A' && arr[level] <= 'Z' {
		dfs(arr, level+1)
		arr[level] = arr[level] - 'A' + 'a' // 小写
		dfs(arr, level+1)
		arr[level] = arr[level] - 'a' + 'A' // 大写
	} else {
		dfs(arr, level+1)
	}
	return
}
