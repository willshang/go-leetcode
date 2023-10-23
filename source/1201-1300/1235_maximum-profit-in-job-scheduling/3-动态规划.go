package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(jobScheduling([]int{1, 2, 3, 4, 6}, []int{3, 5, 10, 6, 9}, []int{20, 20, 100, 70, 60}))
	//fmt.Println(jobScheduling([]int{1, 1, 1}, []int{2, 3, 4}, []int{5, 6, 4}))
}

// leetcode1235_规划兼职工作
type Node struct {
	startTime int
	endTime   int
	profit    int
}

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	n := len(startTime)
	arr := make([]Node, 0)
	for i := 0; i < n; i++ {
		arr = append(arr, Node{
			startTime: startTime[i],
			endTime:   endTime[i],
			profit:    profit[i],
		})
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i].endTime == arr[j].endTime {
			return arr[i].startTime < arr[j].startTime
		}
		return arr[i].endTime < arr[j].endTime
	})
	dp := make([]int, n)
	dp[0] = arr[0].profit
	for i := 1; i < n; i++ {
		dp[i] = dp[i-1]
		dp[i] = max(dp[i], arr[i].profit)
		for j := i - 1; j >= 0; j-- {
			if arr[j].endTime <= arr[i].startTime {
				dp[i] = max(dp[i], dp[j]+arr[i].profit)
				break
			}
		}
	}
	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
