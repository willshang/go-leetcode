package main

func main() {

}

var dp [][]int

func stoneGameII(piles []int) int {
	n := len(piles)
	dp = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := n - 2; i >= 0; i-- {
		piles[i] = piles[i] + piles[i+1]
	}
	return dfs(piles, 0, 1)
}

func dfs(piles []int, index, M int) int {
	if index >= len(piles) {
		return 0
	}
	if index+2*M >= len(piles) { // 可以全部拿走
		return piles[index]
	}
	if dp[index][M] > 0 {
		return dp[index][M]
	}
	res := 0
	for i := 1; i <= 2*M; i++ { // 尝试不同拿法
		res = max(res, piles[index]-dfs(piles, index+i, max(M, i)))
	}
	dp[index][M] = res
	return dp[index][M]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
