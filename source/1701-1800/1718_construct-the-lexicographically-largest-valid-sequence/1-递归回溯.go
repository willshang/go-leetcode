package main

import (
	"fmt"
)

func main() {
	// fmt.Println(constructDistancedSequence(17))
	fmt.Println(constructDistancedSequence(3))
}

// leetcode1718_构建字典序最大的可行序列
var res []int

func constructDistancedSequence(n int) []int {
	path := make([]int, 2*n-1)
	res = make([]int, 2*n-1)
	visited := make([]bool, n+1)
	dfs(n, 0, &path, &visited)
	return res
}

func dfs(n int, cur int, path *[]int, visited *[]bool) bool {
	if cur == 2*n-1 {
		copy(res, *path)
		return true
	}
	if (*path)[cur] > 0 {
		return dfs(n, cur+1, path, visited)
	}
	for i := n; i > 0; i-- { // 从大的开始尝试
		next := cur + i
		if (*visited)[i] == true {
			continue
		}
		if i > 1 && (next >= 2*n-1 || (*path)[next] > 0) {
			continue
		}
		a, b := cur, next
		if i == 1 { // 1特判
			a, b = cur, cur
		}
		(*visited)[i] = true
		(*path)[a] = i
		(*path)[b] = i
		if dfs(n, cur+1, path, visited) == true {
			return true
		}
		(*path)[a] = 0
		(*path)[b] = 0
		(*visited)[i] = false
	}
	return false
}
