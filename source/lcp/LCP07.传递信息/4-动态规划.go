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

// leetcode_lcp07_传递信息
func numWays(n int, relation [][]int, k int) int {
	dp := make([]int, n)
	dp[0] = 1
	for i := 0; i < k; i++ {
		temp := make([]int, n)
		for _, v := range relation {
			temp[v[1]] = temp[v[1]] + dp[v[0]]
		}
		dp = temp
	}
	return dp[n-1]
}
