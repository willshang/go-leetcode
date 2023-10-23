package main

import "fmt"

func main() {
	fmt.Println(combinationSum4([]int{1, 2, 3}, 4))
}

// 剑指OfferII104.排列的数目
func combinationSum4(nums []int, target int) int {
	// 等价于：
	// 假设你正在爬楼梯。需要n阶你才能到达楼顶。
	// 每次你可以爬num(num in nums)级台阶。
	// 你有多少种不同的方法可以爬到楼顶呢？
	dp := make([]int, target+1)
	dp[0] = 1 // 爬0楼1种解法
	for i := 1; i <= target; i++ {
		for j := 0; j < len(nums); j++ {
			if i-nums[j] >= 0 {
				dp[i] = dp[i] + dp[i-nums[j]]
			}
		}
	}
	return dp[target]
}
