package main

import "fmt"

func main() {
	arr := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	fmt.Println(minimumTotal(arr))
}

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := [2][]int{}
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		cur := i % 2
		prev := 1 - cur
		dp[cur][0] = dp[prev][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			dp[cur][j] = min(dp[prev][j-1], dp[prev][j]) + triangle[i][j]
		}
		dp[cur][i] = dp[prev][i-1] + triangle[i][i]
	}
	res := dp[(n-1)%2][0]
	for i := 1; i < n; i++ {
		res = min(res, dp[(n-1)%2][i])
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
