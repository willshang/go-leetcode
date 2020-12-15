package main

import "sort"

func main() {

}

// leetcode1691_堆叠长方体的最大高度
func maxHeight(cuboids [][]int) int {
	for i := 0; i < len(cuboids); i++ {
		arr := cuboids[i]
		sort.Ints(arr)
		cuboids[i] = arr
	}
	sort.Slice(cuboids, func(i, j int) bool {
		if cuboids[i][0] == cuboids[j][0] {
			if cuboids[i][1] == cuboids[j][1] {
				return cuboids[i][2] < cuboids[j][2]
			}
			return cuboids[i][1] < cuboids[j][1]
		}
		return cuboids[i][0] < cuboids[j][0]
	})
	n := len(cuboids)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = cuboids[i][2]
	}
	res := dp[0]
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if cuboids[i][0] >= cuboids[j][0] &&
				cuboids[i][1] >= cuboids[j][1] &&
				cuboids[i][2] >= cuboids[j][2] {
				dp[i] = max(dp[i], dp[j]+cuboids[i][2])
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
