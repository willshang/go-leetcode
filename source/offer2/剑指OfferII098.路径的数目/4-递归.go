package main

import "fmt"

func main() {
	fmt.Println(uniquePaths(3, 2))
}

var arr [][]int

func uniquePaths(m int, n int) int {
	arr = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		arr[i] = make([]int, m+1)
	}
	return dfs(m, n)
}

func dfs(m, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}
	if m == 1 || n == 1 {
		return 1
	}
	if arr[n][m] > 0 {
		return arr[n][m]
	}
	arr[n][m] = dfs(m, n-1) + dfs(m-1, n)
	return arr[n][m]
}
