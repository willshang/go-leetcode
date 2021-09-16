package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minSpaceWastedKResizing([]int{10, 20}, 0))
}

func minSpaceWastedKResizing(nums []int, k int) int {
	n := len(nums)
	arr := make([][]int, n) // arr[i][j]表示nums[i:j]的最小浪费空间
	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		maxValue, sum := math.MinInt32, 0
		for j := i; j < n; j++ {
			if nums[j] > maxValue {
				maxValue = nums[j]
			}
			sum = sum + nums[j]
			arr[i][j] = maxValue*(j-i+1) - sum // 最大值*长度-总和
		}
	}
	dp := make([][]int, n) // dp[i][j]表示将nums[:i]分为j段的最小浪费空间
	for i := 0; i < n; i++ {
		dp[i] = make([]int, k+2)
		for j := 0; j < k+2; j++ {
			dp[i][j] = math.MaxInt32 / 10
		}
	}
	for i := 0; i < n; i++ {
		for j := 1; j <= k+1; j++ { // 调整k次，最少1段，最多k+1段
			for l := 0; l <= i; l++ {
				if l == 0 {
					dp[i][j] = arr[0][i]
				} else {
					dp[i][j] = min(dp[i][j], dp[l-1][j-1]+arr[l][i])
				}
			}
		}
	}
	return dp[n-1][k+1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
