package main

func main() {

}

func maxResult(nums []int, k int) int {
	n := len(nums)
	dp := make([]int, n)
	if k > n {
		k = n
	}
	dp[0] = nums[0]
	maxValue := nums[0]
	for i := 1; i < len(nums); i++ {
		if i <= k { // 在前k个，dp[i] = maxValue + nums[i]
			dp[i] = maxValue + nums[i]
			maxValue = max(maxValue, dp[i])
		} else {
			if maxValue == dp[i-1-k] { // 需要重新找maxValue
				maxValue = getMaxValue(dp[i-k : i])
			}
			dp[i] = maxValue + nums[i]
			maxValue = max(maxValue, dp[i])
		}
	}
	return dp[n-1]
}

func getMaxValue(arr []int) int {
	maxValue := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > maxValue {
			maxValue = arr[i]
		}
	}
	return maxValue
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
