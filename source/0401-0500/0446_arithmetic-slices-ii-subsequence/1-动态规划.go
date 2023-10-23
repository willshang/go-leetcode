package main

func main() {

}

// leetcode446_等差数列划分II-子序列
func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	dp := make([]map[int]int, n) // dp[i][d]代表以A[i]结束且公差为d的等差数列个数
	for i := 0; i < n; i++ {
		dp[i] = make(map[int]int)
	}
	res := 0
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			diff := nums[i] - nums[j]
			dp[i][diff] = dp[i][diff] + dp[j][diff] + 1
			res = res + dp[j][diff]
		}
	}
	return res
}
