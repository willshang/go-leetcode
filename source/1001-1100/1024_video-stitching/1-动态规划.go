package main

import "math"

func main() {

}

func videoStitching(clips [][]int, T int) int {
	dp := make([]int, T+1) //dp[i]表示将区间[0,i)覆盖所需的最少子区间的数量
	for i := 0; i <= T; i++ {
		dp[i] = math.MaxInt32 / 10
	}
	dp[0] = 0
	for i := 1; i <= T; i++ {
		for j := 0; j < len(clips); j++ {
			a, b := clips[j][0], clips[j][1]
			if a < i && i <= b && dp[a]+1 < dp[i] {
				dp[i] = dp[a] + 1
			}
		}
	}
	if dp[T] == math.MaxInt32/10 {
		return -1
	}
	return dp[T]
}
