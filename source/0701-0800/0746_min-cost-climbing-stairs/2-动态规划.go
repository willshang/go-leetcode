package main

import "fmt"

func main() {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	fmt.Println(minCostClimbingStairs(cost))
}

// leetcode746_使用最小花费爬楼梯
func minCostClimbingStairs(cost []int) int {
	a := 0
	b := 0
	for i := 2; i <= len(cost); i++ {
		a, b = b, min(b+cost[i-1], a+cost[i-2])
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
