package main

func main() {

}

var dp []int
var temp []int

func maxScore(nums []int) int {
	n := len(nums)
	total := 1 << n
	dp = make([]int, total)
	temp = make([]int, total)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			status := (1 << i) | (1 << j) // 把第i位和第j位置1
			temp[status] = gcd(nums[i], nums[j])
		}
	}
	return dfs(nums, total-1, 0) // 从目标往前推
}

func dfs(nums []int, status int, count int) int {
	if dp[status] > 0 {
		return dp[status]
	}
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if (status>>i)&1 == 0 || (status>>j)&1 == 0 { // status的第i或者第j位为0
				continue
			}
			cur := (1 << i) | (1 << j)
			prev := status ^ (1 << i) ^ (1 << j) // 异或操作：把status都为1的第i,j位变为0
			dp[status] = max(dp[status], dfs(nums, prev, count+1)+(count+1)*temp[cur])
		}
	}
	return dp[status]
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
