package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxEnvelopes([][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}))
}

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
	dp[0] = 1
	res := 1
	for i := 1; i < len(envelopes); i++ {
		temp := 0
		for j := i - 1; j >= 0; j-- {
			if envelopes[i][0] > envelopes[j][0] &&
				envelopes[i][1] > envelopes[j][1] && dp[j] > temp {
				temp = dp[j]
			}
		}
		dp[i] = temp + 1
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}
