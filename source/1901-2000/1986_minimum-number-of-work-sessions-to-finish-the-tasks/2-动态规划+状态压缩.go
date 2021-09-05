package main

import "fmt"

func main() {
	a := 21
	for i := a; i > 0; i = (i - 1) & a {
		fmt.Printf("%b\n", i)
	}
}

func minSessions(tasks []int, sessionTime int) int {
	n := len(tasks)
	total := 1 << n          // 总的状态数
	dp := make([]int, total) // dp[i] => 任务状态为i时的最少工作时间
	for i := 0; i < total; i++ {
		dp[i] = n // 最多n次
	}
	dp[0] = 0
	valid := make([]bool, total) // 状态是否<=sessionTime
	for i := 1; i < total; i++ { // 枚举状态
		sum := 0 // 该状态和
		for j := 0; j < n; j++ {
			if i&(1<<j) > 0 {
				sum = sum + tasks[j]
			}
		}
		if sum <= sessionTime {
			valid[i] = true
		}
	}

	for i := 1; i < total; i++ { // 枚举状态
		for j := i; j > 0; j = (j - 1) & i { // 遍历得到比较小的子集：数字i二进制为1位置上的非0子集
			if valid[j] == true {
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
