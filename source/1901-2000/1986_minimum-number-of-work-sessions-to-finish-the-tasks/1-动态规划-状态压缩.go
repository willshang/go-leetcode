package main

import "fmt"

func main() {
	//a := 21
	//for i := a; i > 0; i = (i - 1) & a {
	//	fmt.Printf("%b\n", i)
	//}
	fmt.Println(minSessions([]int{1, 1, 1, 1, 1, 1, 1}, 10))
}

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
		for j := i; j > 0; j = (j - 1) & i { // 遍历得到比较小的子集：数字i二进制为1位置上的非0子集
			if sum[j] <= sessionTime {
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
