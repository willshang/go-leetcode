package main

func main() {

}

// leetcode1155_掷骰子的N种方法
func numRollsToTarget(d int, f int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 0; i < d; i++ {
		next := make([]int, target+1)
		for j := 1; j <= f; j++ {
			for k := 0; k <= target-j; k++ {
				next[k+j] = (next[k+j] + dp[k]) % 1000000007
			}
		}
		dp = next
	}
	return dp[target]
}
