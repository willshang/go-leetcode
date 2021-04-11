package main

import "fmt"

func main() {
	arr := [][]int{
		{0, 2},
		{2, 1},
		{3, 4},
		{2, 3},
		{1, 4},
		{2, 0},
		{0, 4},
	}
	fmt.Println(numWays(5, arr, 3))
}

func numWays(n int, relation [][]int, k int) int {
	dp := make([][]int, k+1)
	for i := 0; i < k+1; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	for i := 0; i < k; i++ {
		for _, v := range relation {
			dp[i+1][v[1]] = dp[i+1][v[1]] + dp[i][v[0]]
		}
	}
	return dp[k][n-1]
}
