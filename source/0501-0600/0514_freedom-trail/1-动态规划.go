package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findRotateSteps("edcba", "abcde"))
}

// leetcode514_自由之路
func findRotateSteps(ring string, key string) int {
	maxValue := math.MaxInt32 / 10
	n := len(key)
	m := len(ring)
	dp := make([][]int, n) // dp[i][j] => key[:i+1]，ring[:j+1]的最少步数
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = maxValue
		}
	}
	arr := [26][]int{}
	for i := 0; i < len(ring); i++ {
		value := int(ring[i] - 'a')
		arr[value] = append(arr[value], i)
	}
	for _, v := range arr[key[0]-'a'] {
		dp[0][v] = min(v, m-v) + 1 // 移到次数
	}
	for i := 1; i < n; i++ {
		for _, j := range arr[key[i]-'a'] { // 枚举当前字母位置
			for _, k := range arr[key[i-1]-'a'] { // 枚举上一个字母位置
				minValue := min(abs(j-k), m-abs(j-k))
				dp[i][j] = min(dp[i][j], dp[i-1][k]+minValue+1)
			}
		}
	}
	res := math.MaxInt32
	for i := 0; i < m; i++ {
		res = min(res, dp[n-1][i])
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
