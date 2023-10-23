package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(mincostTickets([]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15}))
}

var dp [366]int
var duration = []int{1, 7, 30}

func mincostTickets(days []int, costs []int) int {
	dp = [366]int{}
	return dfs(0, costs, days)
}

func dfs(day int, costs []int, days []int) int {
	if day >= len(days) {
		return 0
	}
	if dp[day] > 0 {
		return dp[day]
	}
	dp[day] = math.MaxInt32
	j := day
	for i := 0; i < 3; i++ {
		for ; j < len(days) && days[j] < days[day]+duration[i]; j++ {
		}
		dp[day] = min(dp[day], dfs(j, costs, days)+costs[i])
	}
	return dp[day]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
