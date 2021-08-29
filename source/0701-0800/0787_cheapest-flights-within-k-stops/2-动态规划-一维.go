package main

import (
	"fmt"
	"math"
)

func main() {
	arr := [][]int{
		{0, 1, 100},
		{1, 2, 100},
		{0, 2, 500},
	}
	fmt.Println(findCheapestPrice(3, arr, 0, 2, 1))
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	maxValue := math.MaxInt32 / 10
	dp := make([]int, n) // dp[i] => 到j地需要的最少花费（k次中转需要k+1次航班）
	for i := 0; i < n; i++ {
		dp[i] = maxValue
	}
	dp[src] = 0 // 到开始地为0
	res := maxValue
	for i := 1; i <= k+1; i++ {
		temp := make([]int, n)
		for j := 0; j < n; j++ {
			temp[j] = maxValue
		}
		for j := 0; j < len(flights); j++ {
			a, b, c := flights[j][0], flights[j][1], flights[j][2] // a=>b c
			temp[b] = min(temp[b], dp[a]+c)

		}
		res = min(res, temp[dst])
		dp = temp
	}
	if res == maxValue {
		return -1
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
