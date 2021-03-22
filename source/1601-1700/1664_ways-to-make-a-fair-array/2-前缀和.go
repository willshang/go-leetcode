package main

func main() {

}

// leetcode1664_生成平衡数组的方案数
func waysToMakeFair(nums []int) int {
	n := len(nums)
	dp := make([]int, n+1)
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			dp[i+1] = dp[i] + nums[i]
		} else {
			dp[i+1] = dp[i] - nums[i]
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		// dp[i]表示索引i左边部分奇偶元素差值，
		// dp[n] - dp[i+1]表示索引i右边部分奇偶元素差值
		if dp[i] == dp[n]-dp[i+1] {
			res++
		}
	}
	return res
}
