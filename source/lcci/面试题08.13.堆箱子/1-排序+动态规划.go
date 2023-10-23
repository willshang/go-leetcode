package main

import "sort"

func main() {

}

// 程序员面试金典08.13_堆箱子
func pileBox(box [][]int) int {
	sort.Slice(box, func(i, j int) bool {
		if box[i][0] == box[j][0] {
			if box[i][1] == box[j][1] {
				return box[i][2] < box[j][2]
			}
			return box[i][1] < box[j][1]
		}
		return box[i][0] < box[j][0]
	})
	n, res := len(box), 0
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = box[i][2]
	}
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if box[j][0] < box[i][0] && box[j][1] < box[i][1] && box[j][2] < box[i][2] {
				dp[i] = max(dp[i], dp[j]+box[i][2])
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
