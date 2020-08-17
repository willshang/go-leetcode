package main

import "fmt"

func main() {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	fmt.Println(minCostClimbingStairs(cost))
}

var arr []int

func minCostClimbingStairs(cost []int) int {
	arr = make([]int, len(cost)+1)
	return ClimbingStais(cost, len(cost))
}

func ClimbingStais(cost []int, i int) int {
	if i == 0 || i == 1 {
		return 0
	}
	if arr[i] == 0 {
		arr[i] = min(ClimbingStais(cost, i-1)+cost[i-1],
			ClimbingStais(cost, i-2)+cost[i-2])
	}
	return arr[i]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
