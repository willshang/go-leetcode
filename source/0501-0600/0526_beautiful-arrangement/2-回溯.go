package main

import "fmt"

func main() {
	fmt.Println(countArrangement(2))
	fmt.Println(countArrangement(15))
}

// leetcode526_优美的排列
var res int

func countArrangement(n int) int {
	res = 0
	dfs(n, make([]int, 0), make([]bool, n+1))
	return res
}

func dfs(n int, path []int, visited []bool) {
	if len(path) == n {
		res++
		return
	}
	for i := 1; i <= n; i++ {
		index := len(path) + 1
		if visited[i] == false && (i%index == 0 || index%i == 0) {
			visited[i] = true
			dfs(n, append(path, i), visited)
			visited[i] = false
		}
	}
}
