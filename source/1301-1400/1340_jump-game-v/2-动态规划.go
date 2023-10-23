package main

import "sort"

func main() {

}

// leetcode1340_跳跃游戏V
type Node struct {
	index int
	value int
}

func maxJumps(arr []int, d int) int {
	n := len(arr)
	dp := make([]int, n)
	temp := make([]Node, 0)
	for i := 0; i < n; i++ {
		temp = append(temp, Node{index: i, value: arr[i]})
	}
	sort.Slice(temp, func(i, j int) bool {
		if temp[i].value == temp[j].value {
			return temp[i].index < temp[j].index
		}
		return temp[i].value < temp[j].value
	})
	res := 1
	for i := 0; i < n; i++ {
		index := temp[i].index
		dp[index] = 1
		for j := index - 1; j >= 0 && index-j <= d && arr[index] > arr[j]; j-- {
			if dp[j] != 0 {
				dp[index] = max(dp[index], dp[j]+1)
			}
		}
		for j := index + 1; j < len(arr) && j-index <= d && arr[index] > arr[j]; j++ {
			if dp[j] != 0 {
				dp[index] = max(dp[index], dp[j]+1)
			}
		}
		res = max(res, dp[index])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
