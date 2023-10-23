package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(tallestBillboard([]int{1, 2, 3, 6}))
}

const MinValue = math.MinInt32 / 100

// 官方题解(从后往前比较难理解)
// dp[i][s] 表示当我们可以使用 rods[j] (j >= i) 时能得到的最大 score
// dp[i][s] = max(dp[i+1][s], dp[i+1][s-rods[i]], rods[i] + dp[i+1][s+rods[i]])
// 例如:rods=[1,2,3,6],可以有dp[1][1]=5,
// 在写下1之后,可以写下+2,+3,-6使得剩下的rods[i:]获得score为5
var dp [][]int

func tallestBillboard(rods []int) int {
	dp = make([][]int, len(rods))
	for i := 0; i < len(rods); i++ {
		dp[i] = make([]int, 10001)
	}
	res := dfs(rods, 0, 5000)
	return res
}

func dfs(rods []int, index, total int) int {
	if index == len(rods) {
		if total == 5000 {
			return 0
		}
		return MinValue
	}
	if dp[index][total] != 0 {
		return dp[index][total]
	}
	res := dfs(rods, index+1, total)
	res = max(res, dfs(rods, index+1, total-rods[index]))
	res = max(res, dfs(rods, index+1, total+rods[index])+rods[index])
	dp[index][total] = res
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
