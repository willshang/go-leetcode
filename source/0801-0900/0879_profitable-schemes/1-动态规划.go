package main

import "fmt"

func main() {
	fmt.Println(profitableSchemes(5, 3, []int{2, 2}, []int{2, 3}))
}

// leetcode879_盈利计划
var mod = 1000000007

func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
	length := len(group)
	dp := make([][][]int, length+1) // dp[i][j][k] => 在前i个工作，j个员工，利润最少为k 的计划数
	for i := 0; i <= length; i++ {
		dp[i] = make([][]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = make([]int, minProfit+1)
		}
	}
	dp[0][0][0] = 1
	for i := 0; i < length; i++ {
		profitValue := profit[i]
		groupValue := group[i]
		for j := 0; j <= n; j++ {
			for k := 0; k <= minProfit; k++ {
				if j < groupValue { // 人数不够
					dp[i+1][j][k] = dp[i][j][k]
				} else { // 人数足够
					// 利润为负：说明前面j-groupValue个组可以不干活，产生利润为0
					maxValue := max(0, k-profitValue) // 工作利润至少为k,而不是工作利润恰好为k
					dp[i+1][j][k] = (dp[i][j][k] + dp[i][j-groupValue][maxValue]) % mod
				}
			}
		}
	}
	res := 0
	for j := 0; j <= n; j++ {
		res = (res + dp[length][j][minProfit]) % mod
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
