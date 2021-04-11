package main

import "fmt"

func main() {
	fmt.Println(maxJumps([]int{6, 4, 14, 6, 8, 13, 9, 7, 10, 6, 12}, 2))
}

var dp []int

func maxJumps(arr []int, d int) int {
	n := len(arr)
	dp = make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = -1
	}
	for i := 0; i < n; i++ {
		dfs(arr, d, i)
	}
	res := 1
	for i := 0; i < n; i++ {
		res = max(res, dp[i])
	}
	return res
}

func dfs(arr []int, d int, index int) {
	if dp[index] != -1 {
		return
	}
	dp[index] = 1
	for i := index - 1; i >= 0 && index-i <= d && arr[index] > arr[i]; i-- {
		dfs(arr, d, i)
		dp[index] = max(dp[index], dp[i]+1)
	}
	for i := index + 1; i < len(arr) && i-index <= d && arr[index] > arr[i]; i++ {
		dfs(arr, d, i)
		dp[index] = max(dp[index], dp[i]+1)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
