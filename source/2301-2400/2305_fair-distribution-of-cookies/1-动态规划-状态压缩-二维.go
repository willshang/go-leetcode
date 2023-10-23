package main

import "math"

func main() {

}

func distributeCookies(cookies []int, k int) int {
	n := len(cookies)
	total := 1 << n
	sum := make([]int, total)
	for i := 0; i < n; i++ { // 预处理：饼干分配的状态和，分配给某一个人的和
		count := 1 << i
		for j := 0; j < count; j++ {
			sum[count|j] = sum[j] + cookies[i] // 按位或运算：j前面补1=>子集和加上tasks[i]
		}
	}
	dp := make([][]int, k) // f[i][j]=>给前i个人分配工作，饼干的分配情况为j时，最小不公平程度
	for i := 0; i < k; i++ {
		dp[i] = make([]int, total)
	}
	for i := 0; i < total; i++ { // 第0个人的时候
		dp[0][i] = sum[i]
	}
	for i := 1; i < k; i++ {
		for j := 0; j < total; j++ {
			minValue := math.MaxInt32            // dp[i][j]未赋值，为0
			for a := j; a > 0; a = (a - 1) & j { // 遍历得到比较小的子集：数字j二进制为1位置上的非0子集
				// 取子集的补集：j-a 或者 使用异或j^a
				// minValue = min(minValue, max(dp[i-1][j-a], sum[a]))
				minValue = min(minValue, max(dp[i-1][j^a], sum[a]))
			}
			dp[i][j] = minValue
		}
	}
	return dp[k-1][total-1]
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
