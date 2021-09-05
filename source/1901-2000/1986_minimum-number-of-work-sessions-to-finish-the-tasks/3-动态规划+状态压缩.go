package main

import "fmt"

func main() {
	a := 21
	for i := a; i > 0; i = (i - 1) & a {
		fmt.Printf("%b\n", i)
	}
}

// leetcode1986_完成任务的最少工作时间段
func minSessions(tasks []int, sessionTime int) int {
	n := len(tasks)
	total := 1 << n          // 总的状态数
	dp := make([]int, total) // dp[i] => 任务状态为i时的最少工作时间
	for i := 0; i < total; i++ {
		dp[i] = n // 最多n次
	}
	dp[0] = 0
	sum := make([]int, total) // 枚举任务所有状态的和
	for i := 0; i < n; i++ {
		count := 1 << i
		for j := 0; j < count; j++ {
			sum[count|j] = sum[j] + tasks[i] // 按位或运算：j前面补1=>子集和加上tasks[i]
		}
	}
	for i := 0; i < total; i++ {
		for j := 1; j <= i; j++ { // 暴力枚举子集
			if i|j == i && sum[j] <= sessionTime {
				dp[i] = min(dp[i], dp[i^j]+1) // 取补集-异或操作：取dp[i^j]+1操作最小值
			}
		}
	}
	return dp[total-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
