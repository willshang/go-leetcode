package main

func main() {

}

// leetcode1049_最后一块石头的重量II
func lastStoneWeightII(stones []int) int {
	sum := 0
	for i := 0; i < len(stones); i++ {
		sum = sum + stones[i]
	}
	// 求最后1个最小，把石头分2堆，求差值
	// 题目转换为0-1背包问题，容量为sum/2，能装多大体积
	target := sum / 2
	dp := make([]int, sum/2+1)
	for i := 0; i < len(stones); i++ {
		for j := target; j >= 0; j-- {
			if j-stones[i] >= 0 {
				dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
			}
		}
	}
	return sum - 2*dp[sum/2]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
