package main

import "fmt"

func main() {
	fmt.Println(tallestBillboard([]int{1, 2, 3, 6}))
}

func tallestBillboard(rods []int) int {
	dp := make(map[int]int) // 保存高度差为i时最低一边的高度
	dp[0] = 0
	// 每个rods[i]有3个选择，+rods[i]、-rods[i]、0
	for i := 0; i < len(rods); i++ {
		temp := make(map[int]int)
		for k, v := range dp {
			temp[k] = v
		}
		for k, v := range temp {
			dp[k+rods[i]] = max(dp[k+rods[i]], v)                           // 往高的一侧加，高度差变大，但最低一边高度不变
			dp[k] = dp[k]                                                   // 不加，高度差不变，可忽略
			dp[abs(k-rods[i])] = max(dp[abs(k-rods[i])], v+min(k, rods[i])) // 往低的一侧加，高度差有变化，高度差有增长
		}
	}
	return dp[0]
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
