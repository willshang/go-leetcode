package main

func main() {

}

func constrainedSubsetSum(nums []int, k int) int {
	n := len(nums)
	dp := make([]int, n)
	if k > n {
		k = n
	}
	res := nums[0]
	dp[0] = nums[0]
	maxValue := nums[0]
	for i := 1; i < len(nums); i++ {
		if i <= k { // 在前k个，dp[i] = maxValue + nums[i]
			if maxValue < 0 {
				dp[i] = nums[i]
			} else {
				dp[i] = maxValue + nums[i]
			}
			maxValue = max(maxValue, dp[i])
		} else {
			if maxValue == dp[i-1-k] { // 需要重新找maxValue
				maxValue = getMaxValue(dp[i-k : i])
			}
			if maxValue < 0 {
				dp[i] = nums[i]
			} else {
				dp[i] = maxValue + nums[i]
			}
			maxValue = max(maxValue, dp[i])
		}
		res = max(res, dp[i])
	}
	return res
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
