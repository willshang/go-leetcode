package main

import "fmt"

func main() {
	fmt.Println(maxResult([]int{1, -1, -2, 4, -7, 3}, 2))
}

func maxResult(nums []int, k int) int {
	n := len(nums)
	dp := make([]int, n)
	if k > n {
		k = n
	}
	dp[0] = nums[0]
	maxValue := nums[0]
	maxIndex := 0
	for i := 1; i < len(nums); i++ {
		if i <= k { // 在前k个，dp[i] = maxValue + nums[i]
			dp[i] = maxValue + nums[i]
			if dp[i] >= maxValue {
				maxValue = dp[i]
				maxIndex = i
			}
		} else {
			if i-k > maxIndex {
				maxValue = dp[maxIndex+1]
				for j := maxIndex + 1; j < i; j++ {
					if dp[j] >= maxValue {
						maxValue = dp[j]
						maxIndex = j
					}
				}
			}
			dp[i] = maxValue + nums[i]
			if dp[i] >= maxValue {
				maxValue = dp[i]
				maxIndex = i
			}
		}
	}
	return dp[n-1]
}
