package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(tallestBillboard([]int{1, 2, 3, 6}))
}

// dp[i][s] 表示当我们可以使用 rods[j] (j >= i) 时能得到的最大 score
// dp[i][s] = max(dp[i+1][s], dp[i+1][s-rods[i]], rods[i] + dp[i+1][s+rods[i]])
// 例如:rods=[1,2,3,6],可以有dp[1][1]=5,
// 在写下1之后,可以写下+2,+3,-6使得剩下的rods[i:]获得score为5
const MinValue = math.MinInt32 / 100

func tallestBillboard(rods []int) int {
	dp := make([][]int, len(rods)+1)
	for i := 0; i < len(rods)+1; i++ {
		dp[i] = make([]int, 10001)
	}
	for j := 0; j < len(dp[len(rods)]); j++ {
		dp[len(rods)][j] = MinValue
	}
	sum := 0
	for i := 0; i < len(rods); i++ {
		sum = sum + rods[i]
	}
	m := 2 * sum
	dp[len(rods)][sum] = 0
	for i := len(rods) - 1; i >= 0; i-- {
		for s := rods[i]; s <= m-rods[i]; s++ {
			dp[i][s] = max(dp[i+1][s], max(dp[i+1][s-rods[i]], rods[i]+dp[i+1][s+rods[i]]))
		}
	}
	return dp[0][sum]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
