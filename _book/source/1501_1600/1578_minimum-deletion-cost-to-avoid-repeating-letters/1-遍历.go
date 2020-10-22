package main

import "fmt"

func main() {
	fmt.Println(minCost("aaaaa", []int{6, 8, 3, 4, 5}))
}

func minCost(s string, cost []int) int {
	res := 0
	cur := 0
	for i := 0; i < len(s)-1; i++ {
		if s[cur] == s[i+1] {
			res = res + min(cost[cur], cost[i+1])
			if cost[cur] < cost[i+1] {
				cur = i + 1 // 相同，保存较大的
			}
		} else {
			cur = i + 1
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
