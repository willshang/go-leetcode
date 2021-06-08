package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(shortestSubarray([]int{2, -1, 2, 3, 4, -5, 1, 2, 3}, 3))
}

// leetcode862_和至少为K的最短子数组
func shortestSubarray(nums []int, k int) int {
	res := math.MaxInt32
	n := len(nums)
	arr := make([]int, n+1)
	for i := 1; i <= n; i++ {
		arr[i] = arr[i-1] + nums[i-1]
	}
	queue := make([]int, 0) // 递增队列：队首坐标小，队尾坐标大
	for i := 0; i <= n; i++ {
		for len(queue) > 0 && arr[i] <= arr[queue[len(queue)-1]] { // 前缀和小于队尾值：出队
			queue = queue[:len(queue)-1]
		}
		for len(queue) > 0 && arr[i]-arr[queue[0]] >= k { // 差值大于等于k
			res = min(res, i-queue[0])
			queue = queue[1:]
		}
		queue = append(queue, i)
	}
	if res == math.MaxInt32 {
		return -1
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
