package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(jobScheduling([]int{1, 2, 3, 4, 6}, []int{3, 5, 10, 6, 9}, []int{20, 20, 100, 70, 60}))
	//fmt.Println(jobScheduling([]int{1, 1, 1}, []int{2, 3, 4}, []int{5, 6, 4}))
}

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
		left, right := 0, i-1
		for left < right {
			mid := left + (right-left)/2
			if arr[mid+1].endTime <= arr[i].startTime {
				left = mid + 1
			} else {
				right = mid
			}
		}
		if arr[left].endTime <= arr[i].startTime {
			dp[i] = max(dp[i-1], dp[left]+arr[i].profit)
		} else {
			dp[i] = max(dp[i-1], arr[i].profit)
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
