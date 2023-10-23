package main

import "fmt"

func main() {
	fmt.Println(PredictTheWinner([]int{1, 5, 2}))
}

// leetcode486_预测赢家
func PredictTheWinner(nums []int) bool {
	return dfs(nums, 0, len(nums)-1) >= 0
}

func dfs(nums []int, start, end int) int {
	if start > end {
		return 0
	}
	// 玩家得分：自己得分-对手得分
	left := nums[start] - dfs(nums, start+1, end)
	right := nums[end] - dfs(nums, start, end-1)
	return max(left, right)

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
