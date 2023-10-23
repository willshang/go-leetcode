package main

import "math"

func main() {

}

// leetcode2017_网格游戏
func gridGame(grid [][]int) int64 {
	n := len(grid[0])
	a := make([]int, n) // 前缀和：上边：从右到左
	b := make([]int, n) // 前缀和：下边，从左到右
	for i := n - 1; i > 0; i-- {
		a[i-1] = a[i] + grid[0][i]
	}
	for i := 0; i < n-1; i++ {
		b[i+1] = b[i] + grid[1][i]
	}
	res := math.MaxInt64
	// 当第一个机器人选择第i点往下走的时候
	// 第一个机器人不需要选择最大值，只需要考虑让第二个机器人选择最小
	for i := 0; i < n; i++ {
		// 第二个机器人只有2个选择，选其中最大的
		// 1、从第0个点往下走，拿到b[i]值
		// 2、一直往右走，拿到a[i]值
		res = min(res, max(a[i], b[i]))
	}
	return int64(res)
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
