package main

import "math/bits"

func main() {

}

// leetcode1947_最大兼容性评分和
func maxCompatibilitySum(students [][]int, mentors [][]int) int {
	n := len(students)
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
		for j := 0; j < n; j++ {
			arr[i][j] = calculate(students[i], mentors[j])
		}
	}
	target := 1 << n
	dp := make([]int, target) // 所有的状态
	for i := 1; i < target; i++ {
		count := bits.OnesCount(uint(i))
		for j := 0; j < n; j++ {
			if i&(1<<j) != 0 { // 判断第j位是否为1
				prev := i ^ (1 << j)
				dp[i] = max(dp[i], dp[prev]+arr[count-1][j])
			}
		}
	}
	return dp[target-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func calculate(a, b []int) int {
	res := 0
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			res++
		}
	}
	return res
}
