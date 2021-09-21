package main

import "sort"

func main() {

}

func maxTaxiEarnings(n int, rides [][]int) int64 {
	m := len(rides)
	arr := make([][]int64, 0)
	for i := 0; i < m; i++ {
		a, b, c := rides[i][0], rides[i][1], rides[i][2]
		arr = append(arr, []int64{int64(a), int64(b), int64(b - a + c)})
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i][1] == arr[j][1] {
			return arr[i][0] < arr[j][0]
		}
		return arr[i][1] < arr[j][1]
	})
	dp := make([]int64, m)
	dp[0] = arr[0][2]
	for i := 1; i < m; i++ {
		left, right := 0, i-1
		for left < right {
			mid := left + (right-left)/2
			if arr[mid+1][1] <= arr[i][0] {
				left = mid + 1
			} else {
				right = mid
			}
		}
		if arr[left][1] <= arr[i][0] {
			dp[i] = max(dp[i-1], dp[left]+arr[i][2])
		} else {
			dp[i] = max(dp[i-1], arr[i][2])
		}
	}
	return dp[m-1]
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
