package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(largestNumber([]int{4, 3, 2, 5, 6, 7, 2, 5, 5}, 8))
}

func largestNumber(cost []int, target int) string {
	dp, path := make([][]int, 10), make([][]int, 10)
	for i := 0; i < 10; i++ {
		dp[i], path[i] = make([]int, target+1), make([]int, target+1)
		for j := 0; j <= target; j++ {
			dp[i][j] = math.MinInt32
		}
	}
	dp[0][0] = 0 // dp[i+1][j] 表示使用前i个数且花费总代价恰好为j时的最大长度
	for i := 0; i < len(cost); i++ {
		for j := 0; j <= target; j++ {
			if j < cost[i] { // 不够
				dp[i+1][j] = dp[i][j]
				path[i+1][j] = j
			} else {
				// 不选:dp[i][j]、选:dp[i+1][j-cost[i]]
				if dp[i][j] > dp[i+1][j-cost[i]]+1 { // 不选
					dp[i+1][j] = dp[i][j]
					path[i+1][j] = j
				} else { // 选：相等时，也要更新，cost[i]相同取更大i
					dp[i+1][j] = dp[i+1][j-cost[i]] + 1
					path[i+1][j] = j - cost[i]
				}
			}
		}
	}
	if dp[9][target] < 0 {
		return "0"
	}
	res := ""
	i, j := 9, target
	for i > 0 {
		if j == path[i][j] {
			i--
		} else {
			res = res + string('0'+i)
			j = path[i][j]
		}
	}
	return res
}
