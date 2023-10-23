package main

import (
	"fmt"
	"math"
	"sort"
)

var maxConstValue = math.MaxInt32 / 10

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		var tempA int
		fmt.Scan(&tempA)
		a[i] = tempA
	}
	var m int
	fmt.Scan(&m)
	b, c := make([]int, m), make([]int, m)
	for i := 0; i < m; i++ {
		var tempB, tempC int
		fmt.Scan(&tempB, &tempC)
		b[i] = tempB
		c[i] = tempC
	}
	res := getResult(a, b, c)
	fmt.Println(res)
}

func getResult(a, b, c []int) string {
	sort.Ints(a)
	maxValue := a[len(a)-1]
	// dp[i] => 构造长度为i的诗歌长度耗时
	dp := make([]int, maxValue+1)
	for i := 0; i <= maxValue; i++ {
		dp[i] = maxConstValue
	}
	dp[0] = 0
	// 对应的诗歌b[i]需要c[i]的时间,其中b中的数各不相同
	for i := 0; i < len(b); i++ {
		if b[i] <= maxValue {
			dp[b[i]] = c[i]
		}
	}
	for i := 1; i <= maxValue; i++ {
		for j := 0; j < len(b); j++ {
			if b[j] <= i { // 长度i的耗时=长度j的耗时+长度i-j的耗时，找到最小的
				target := i - b[j]
				dp[i] = min(dp[i], dp[target]+c[j])
			}
		}
	}
	res := maxConstValue
	for i := 0; i < len(a); i++ {
		res = min(res, dp[a[i]])
	}
	if res == maxConstValue {
		return "It is not true!"
	}
	return fmt.Sprintf("%d", res)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
