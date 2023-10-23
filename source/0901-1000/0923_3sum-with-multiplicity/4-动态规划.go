package main

import (
	"fmt"
)

func main() {
	fmt.Println(threeSumMulti([]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}, 8))
}

func threeSumMulti(arr []int, target int) int {
	n := len(arr)
	dp := make([][4][]int, n+1) // 虑前i个数时，从中选出j个数，组成k大小的方案数
	for i := 0; i <= n; i++ {
		dp[i] = [4][]int{}
		for j := 0; j < 4; j++ {
			dp[i][j] = make([]int, target+1)
		}
	}
	for i := 0; i <= n; i++ {
		dp[i][0][0] = 1
	}
	for i := 1; i <= n; i++ {
		for j := 1; j < 4; j++ {
			for k := 0; k <= target; k++ {
				if k >= arr[i-1] { // 可选
					dp[i][j][k] = (dp[i][j][k] + dp[i-1][j-1][k-arr[i-1]]) % 1000000007
				}
				dp[i][j][k] = (dp[i][j][k] + dp[i-1][j][k]) % 1000000007 // 不选
			}
		}
	}
	return dp[n][3][target]
}
