package main

import "fmt"

func main() {
	fmt.Println(numOfSubarrays([]int{1, 3, 5}))
}

// leetcode1524_和为奇数的子数组数目
func numOfSubarrays(arr []int) int {
	res := 0
	dp := make([][2]int, len(arr))
	// dp[i][0]子数组和为奇数的个数
	// dp[i][1]子数组和为偶数的个数
	if arr[0]%2 == 1 {
		dp[0][0] = 1
		res++
	} else {
		dp[0][1] = 1
	}
	for i := 1; i < len(arr); i++ {
		if arr[i]%2 == 1 {
			dp[i][0] = dp[i-1][1] + 1
			dp[i][1] = dp[i-1][0]
		} else {
			dp[i][0] = dp[i-1][0]
			dp[i][1] = dp[i-1][1] + 1
		}
		res = res + dp[i][0]
	}
	return res % 1000000007
}
