package main

import "fmt"

func main() {
	fmt.Println(mincostTickets([]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15}))
}

var dp [366]int
var m map[int]bool

func mincostTickets(days []int, costs []int) int {
	dp = [366]int{}
	m = make(map[int]bool)
	for i := 0; i < len(days); i++ {
		m[days[i]] = true
	}
	return dfs(1, costs)
}

func dfs(day int, costs []int) int {
	if day > 365 {
		return 0
	}
	if dp[day] > 0 {
		return dp[day]
	}
	if m[day] == true {
		dp[day] = min(min(dfs(day+1, costs)+costs[0], dfs(day+7, costs)+costs[1]),
			dfs(day+30, costs)+costs[2])
	} else {
		dp[day] = dfs(day+1, costs)
	}
	return dp[day]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
