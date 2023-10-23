package main

import "math"

func main() {

}

// leetcode1884_鸡蛋掉落-两枚鸡蛋
func twoEggDrop(n int) int {
	dp := make([]int, n+1) // dp[j]验证j层楼需要的最少操作次数
	for j := 1; j < n+1; j++ {
		dp[j] = math.MaxInt32
	}
	// 值为0
	// dp[0] = 0
	for j := 1; j <= n; j++ {
		for k := 1; k <= j; k++ {
			a := k           //  假设在k层破碎：剩下1枚鸡蛋,求k-1层结果，(k-1)+1
			b := dp[j-k] + 1 //  假设在k层没有破碎：剩余2枚鸡蛋求j-k层结果
			dp[j] = min(dp[j], max(a, b))
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
