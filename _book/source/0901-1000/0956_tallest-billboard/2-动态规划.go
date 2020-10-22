package main

import "fmt"

func main() {
	fmt.Println(tallestBillboard([]int{1, 2, 3, 6}))
}

// leetcode956_最高的广告牌
func tallestBillboard(rods []int) int {
	dp := make(map[int]int) // 保存和为i中正整数的和
	dp[0] = 0
	// 每个rods[i]有3个选择，+rods[i]、-rods[i]、0
	for i := 0; i < len(rods); i++ {
		value := rods[i]
		temp := make(map[int]int)
		for k, v := range dp {
			temp[k+value] = max(temp[k+value], v+value) // +value>0，是正整数，需要加上
			temp[k] = max(temp[k], v)
			temp[k-value] = max(temp[k-value], v) // -value<0，不需要加上
		}
		dp = temp
	}
	return dp[0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
