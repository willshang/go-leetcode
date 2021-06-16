package main

import "fmt"

func main() {
	fmt.Println(maxSumMinProduct([]int{1, 2, 3, 2}))
}

// leetcode1856_子数组最小乘积的最大值
var mod = 1000000007

func maxSumMinProduct(nums []int) int {
	res := 0
	n := len(nums)
	arr := make([]int, n+1) // 前缀和
	for i := 1; i <= n; i++ {
		arr[i] = arr[i-1] + nums[i-1]
	}
	left, right := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		left[i] = -1 // 默认是最左边
		right[i] = n // 默认是最右边
	}
	stack := make([]int, 0) // 双栈：leetcode 84.柱状图中最大的矩形
	for i := 0; i < n; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] > nums[i] {
			right[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	stack = make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] > nums[i] {
			left[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	for i := 0; i < n; i++ {
		target := (arr[right[i]] - arr[left[i]+1]) * nums[i]
		res = max(res, target)
	}
	return res % mod
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
