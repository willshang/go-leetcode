package main

import "fmt"

func main() {
	fmt.Println(maxSubarraySumCircular([]int{-2}))
	fmt.Println(maxSubarraySumCircular([]int{-3, -2, -3}))
}

// leetcode918_环形子数组的最大和
func maxSubarraySumCircular(A []int) int {
	n := len(A)
	// leetcode53.最大子序和 的变形
	total := A[0]                  // 总和
	minValue, sumMin := A[0], A[0] // 找到连续数组之和的最小值
	maxValue, sumMax := A[0], A[0] // 找到连续数组之和的最大值
	for i := 1; i < n; i++ {
		total = total + A[i]
		sumMin = min(sumMin+A[i], A[i])
		minValue = min(minValue, sumMin)
		sumMax = max(sumMax+A[i], A[i])
		maxValue = max(maxValue, sumMax)
	}
	// 2种情况：取最大值即可
	// 1：目标数组不需要环，求最大值
	// 2: 需要成环，总和减去最小值
	if total == minValue { // 极端情况全小于0：如果和=最小值，返回最大值
		return maxValue
	}
	return max(maxValue, total-minValue)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
