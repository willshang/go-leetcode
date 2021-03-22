package main

import "fmt"

func main() {
	fmt.Println(minCost("aaaaa", []int{6, 8, 3, 4, 5}))
}

// leetcode1578_避免重复字母的最小删除成本
func minCost(s string, cost []int) int {
	res := 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			res = res + min(cost[i], cost[i+1])
			if cost[i] > cost[i+1] {
				cost[i], cost[i+1] = cost[i+1], cost[i] // 相同，把较大的存在后面
			}
		}
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
