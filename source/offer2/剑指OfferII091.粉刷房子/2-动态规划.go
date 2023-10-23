package main

import "fmt"

func main() {
	fmt.Println(minCost([][]int{
		{17, 2, 17},
		{16, 16, 5},
		{14, 3, 19},
	}))
}

func minCost(costs [][]int) int {
	n := len(costs)
	a, b, c := costs[0][0], costs[0][1], costs[0][2]
	for i := 1; i < n; i++ {
		a, b, c = min(b, c)+costs[i][0], min(a, c)+costs[i][1], min(a, b)+costs[i][2]
	}
	return min(a, min(b, c))
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
