package main

import "sort"

func main() {

}

// leetcode1626_无矛盾的最佳球队
func bestTeamScore(scores []int, ages []int) int {
	arr := make([][2]int, 0)
	for i := 0; i < len(ages); i++ {
		arr = append(arr, [2]int{ages[i], scores[i]})
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i][0] == arr[j][0] {
			return arr[i][1] < arr[j][1]
		}
		return arr[i][0] < arr[j][0]
	})
	dp := make([]int, len(arr))
	res := 0
	for i := 0; i < len(arr); i++ {
		dp[i] = arr[i][1]
		for j := 0; j < i; j++ {
			if arr[j][1] <= arr[i][1] {
				dp[i] = max(dp[i], dp[j]+arr[i][1])
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
