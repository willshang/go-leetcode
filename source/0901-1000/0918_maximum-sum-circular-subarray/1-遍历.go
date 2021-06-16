package main

import "fmt"

func main() {
	fmt.Println(maxSubarraySumCircular([]int{-2}))
}

func maxSubarraySumCircular(A []int) int {
	n := len(A)
	if n == 1{
		return A[0]
	}
	total := 0 // 总和
	for i := 0; i < n; i++ {
		total = total + A[i]
	}
	minValue := A[0] // 找到连续数组之和的最小值
	sum := 0
	for i := 1; i < n; i++ {
		if sum < 0 {
			sum = sum + A[i]
		} else {
			sum = A[i]
		}
		minValue = min(minValue, sum)
	}
	maxValue := 0
	sum = 0
	for i := 1; i < n; i++ {
		if sum > 0 {
			sum = sum + A[i]
		} else {
			sum = A[i]
		}
		maxValue = max(maxValue, sum)
	}
	// 2种情况：取最大值即可
	// 1：目标数组不需要环，求最大值
	// 2: 需要成环，总和减去最小值
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

