package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(largestSumOfAverages([]int{9, 1, 2, 3, 9}, 3))
}

// leetcode813_最大平均值和的分组
func largestSumOfAverages(A []int, K int) float64 {
	n := len(A)
	arr := make([]int, n+1)
	for i := 0; i < n; i++ {
		arr[i+1] = arr[i] + A[i]
	}
	dp := make([][]float64, n) // dp[i]=>A[i:]的最大平均值
	for i := 0; i < n; i++ {
		dp[i] = make([]float64, K)
		dp[i][0] = float64(arr[n]-arr[i]) / float64(n-i) // 划分为1组
	}
	for k := 1; k < K; k++ { // K组可以划分K-1次
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				target := dp[j][k-1] + float64(arr[j]-arr[i])/float64(j-i)
				dp[i][k] = math.Max(dp[i][k], target)
			}
		}
	}
	return dp[0][K-1]
}
