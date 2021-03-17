package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(findRotateSteps("edcba", "abcde"))
	fmt.Println(findRotateSteps("godding", "gd"))
}

var dp [][]int
var arr [26][]int

func findRotateSteps(ring string, key string) int {
	n := len(key)
	m := len(ring)
	dp = make([][]int, n) // dp[i][j] => key[:i+1]，ring[:j+1]的最少步数
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = -1
		}
	}
	arr = [26][]int{}
	for i := 0; i < len(ring); i++ {
		value := int(ring[i] - 'a')
		arr[value] = append(arr[value], i)
	}
	return n + dfs(key, ring, 0, 0)
}

func dfs(key, ring string, keyIndex, ringIndex int) int {
	if keyIndex == len(key) {
		return 0
	}
	if dp[keyIndex][ringIndex] != -1 {
		return dp[keyIndex][ringIndex]
	}
	cur := int(key[keyIndex] - 'a')
	res := math.MaxInt32
	for _, v := range arr[cur] {
		minValue := min(abs(ringIndex-v), len(ring)-abs(ringIndex-v))
		res = min(res, minValue+dfs(key, ring, keyIndex+1, v))
	}
	dp[keyIndex][ringIndex] = res
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
