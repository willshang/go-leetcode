package main

import "math"

func main() {

}

// leetcode1326_灌溉花园的最少水龙头数目
func minTaps(n int, ranges []int) int {
	// 处理成一组start-end的数组
	// 题目变成 leetcode 1024.视频拼接
	arr := make([][2]int, n+1)
	for i := 0; i <= n; i++ {
		l := max(0, i-ranges[i])
		r := min(n, i+ranges[i])
		arr = append(arr, [2]int{l, r})
	}
	dp := make([]int, n+1) //dp[i]表示将区间[0,i)覆盖所需的最少子区间的数量
	for i := 0; i <= n; i++ {
		dp[i] = math.MaxInt32 / 10
	}
	dp[0] = 0
	for i := 1; i <= n; i++ {
		for j := 0; j < len(arr); j++ {
			a, b := arr[j][0], arr[j][1]
			if a < i && i <= b && dp[a]+1 < dp[i] {
				dp[i] = dp[a] + 1
			}
		}
	}
	if dp[n] == math.MaxInt32/10 {
		return -1
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
