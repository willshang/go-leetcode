package main

func main() {

}

func maxAscendingSum(nums []int) int {
	res := nums[0]
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		res = max(res, dp[i])
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
