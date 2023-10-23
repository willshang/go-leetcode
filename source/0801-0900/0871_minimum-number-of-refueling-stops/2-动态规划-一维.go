package main

func main() {

}

// leetcode871_最低加油次数
func minRefuelStops(target int, startFuel int, stations [][]int) int {
	n := len(stations)
	dp := make([]int, n+1)
	dp[0] = startFuel
	for i := 0; i < n; i++ {
		for j := i; j >= 0; j-- {
			if dp[j] >= stations[i][0] {
				dp[j+1] = max(dp[j+1], dp[j]+stations[i][1])
			}
		}
	}
	for i := 0; i <= n; i++ {
		if dp[i] >= target {
			return i
		}
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
