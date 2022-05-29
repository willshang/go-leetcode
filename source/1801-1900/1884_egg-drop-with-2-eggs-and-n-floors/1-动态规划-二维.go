package main

import "math"

func main() {

}

func twoEggDrop(n int) int {
	dp := [3][]int{} // dp[i][j] 有i枚鸡蛋，验证j层楼需要的最少操作次数
	for i := 0; i < 3; i++ {
		dp[i] = make([]int, n+1)
		for j := 1; j < n+1; j++ {
			dp[i][j] = math.MaxInt32
			if i == 1 {
				dp[1][j] = j // 1个鸡蛋，需要j次
			}
		}
	}
	// 值为0
	// dp[1][0], dp[2][0] = 0,0
	for j := 1; j <= n; j++ {
		for k := 1; k <= j; k++ { //
			// a := k              //  假设在k层破碎：剩下1枚鸡蛋,求k-1层结果，(k-1)+1
			a := dp[1][k-1] + 1 //  假设在k层破碎。剩下1枚鸡蛋 求k-1层结果
			b := dp[2][j-k] + 1 //  假设在k层没有破碎：剩余2枚鸡蛋求j-k层结果
			dp[2][j] = min(dp[2][j], max(a, b))
		}
	}
	return dp[2][n]
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
