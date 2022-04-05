package main

import (
	"fmt"
	"math"
)

func main() {
	arr := [][]int{
		{2, 3},
		{3, 4},
	}
	fmt.Println(minimumFinishTime(arr, 5, 4))
}

// leetcode2188_完成比赛的最少时间
func minimumFinishTime(tires [][]int, changeTime int, numLaps int) int {
	minArr := make([]int, 20) // 用1个轮胎跑i圈的最小花费
	for i := 0; i < 20; i++ {
		minArr[i] = math.MaxInt32 / 10
	}
	for i := 0; i < len(tires); i++ {
		f, r := tires[i][0], tires[i][1]
		first := f + changeTime // 第一次用该胎的耗费
		sum := first
		for j := 1; f <= first; j++ {
			minArr[j] = min(minArr[j], sum)
			f = f * r
			sum = sum + f
		}
	}
	dp := make([]int, numLaps+1) // dp[i] => 表示跑i圈的最小花费
	for i := 0; i <= numLaps; i++ {
		dp[i] = math.MaxInt32 / 10
	}
	dp[0] = 0
	for i := 1; i <= numLaps; i++ {
		for j := 1; j <= 19 && j <= i; j++ { // 遍历跑i圈换胎
			dp[i] = min(dp[i], dp[i-j]+minArr[j])
		}
	}
	return dp[numLaps] - changeTime // 减去最后一次换胎时间
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
