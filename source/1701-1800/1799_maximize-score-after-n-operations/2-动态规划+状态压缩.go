package main

import "math/bits"

func main() {

}

// leetcode1799_N次操作后的最大分数和
func maxScore(nums []int) int {
	n := len(nums)
	total := 1 << n
	dp := make([]int, total)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			status := (1 << i) | (1 << j) // 把第i位和第j位置1
			dp[status] = gcd(nums[i], nums[j])
		}
	}
	for i := 0; i < total; i++ { // 枚举状态：当前状态依赖之前的
		a := bits.OnesCount(uint(i))
		if a%2 == 1 { // 可去除；剪枝：1的个数为奇数个
			continue
		}
		for j := i; j > 0; j = (j - 1) & i { // 遍历i二进制中位置1的子集
			b := bits.OnesCount(uint(j))
			if a-b == 2 { // 1的个数相差2
				dp[i] = max(dp[i], dp[j]+a/2*dp[i^j]) // 取补集-异或操作：i^j
			}
		}
	}
	return dp[total-1]
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
