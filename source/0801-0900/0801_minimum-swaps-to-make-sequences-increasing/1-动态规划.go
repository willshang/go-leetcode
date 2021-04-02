package main

func main() {

}

// leetcode801_使序列递增的最小交换次数
func minSwap(A []int, B []int) int {
	n := len(A)
	dp := make([][2]int, n)
	dp[0][0] = 0 // dp[i][0] 第i个位置不换
	dp[0][1] = 1 // dp[i][1] 第i个位置换
	for i := 1; i < n; i++ {
		if A[i-1] < A[i] && B[i-1] < B[i] {
			if A[i-1] < B[i] && B[i-1] < A[i] { // 可换可不换
				dp[i][0] = min(dp[i-1][0], dp[i-1][1])
				dp[i][1] = min(dp[i-1][0], dp[i-1][1]) + 1
			} else {
				dp[i][0] = dp[i-1][0]     // 不交换则上一轮也不交换
				dp[i][1] = dp[i-1][1] + 1 // 交换则上一轮也交换
			}
		} else {
			dp[i][0] = dp[i-1][1]     // 不交换则上一轮必须交换
			dp[i][1] = dp[i-1][0] + 1 // 交换，则上一轮不能交换
		}
	}
	return min(dp[n-1][0], dp[n-1][1])
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
