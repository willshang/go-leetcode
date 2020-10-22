package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxEnvelopes([][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}))
}

// leetcode354_俄罗斯套娃信封问题
func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) <= 1 {
		return len(envelopes)
	}
	// 宽[0] 高[1]排序
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] < envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})
	// 第i个信封套几个
	dp := make([]int, len(envelopes))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	res := 1
	for i := 1; i < len(envelopes); i++ {
		for j := 0; j < i; j++ {
			if envelopes[i][0] > envelopes[j][0] &&
				envelopes[i][1] > envelopes[j][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
