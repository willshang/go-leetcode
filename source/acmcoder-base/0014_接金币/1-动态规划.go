package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([][2]int, 0)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		arr = append(arr, [2]int{a, b})
	}
	res := getResult(arr)
	fmt.Println(res)
}

func getResult(arr [][2]int) int {
	maxValue := 0
	n := len(arr)
	for i := 0; i < n; i++ {
		maxValue = max(maxValue, arr[i][0])
	}
	// dp[i][j] => 最后在时间i位置j 能获取到的金币数
	dp := make([][11]int, maxValue+1)
	for i := 0; i < n; i++ {
		a, b := arr[i][0], arr[i][1]
		dp[b][a]++
	}
	// 倒推
	for i := maxValue - 1; i >= 0; i-- {
		for j := 0; j < 11; j++ {
			if j == 0 {
				dp[i][j] = dp[i][j] + max(dp[i+1][j], dp[i+1][j+1])
			} else if j == 10 {
				dp[i][j] = dp[i][j] + max(dp[i+1][j], dp[i+1][j-1])
			} else {
				dp[i][j] = dp[i][j] + max(dp[i+1][j], max(dp[i+1][j-1], dp[i+1][j+1]))
			}
		}
	}
	return dp[0][5]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
