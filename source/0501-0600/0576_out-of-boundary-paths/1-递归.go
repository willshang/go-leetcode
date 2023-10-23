package main

func main() {

}

var dp [60][60][60]int
var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}
var mod = 1000000007

func findPaths(m int, n int, N int, i int, j int) int {
	dp = [60][60][60]int{}
	for i := 0; i < 60; i++ {
		for j := 0; j < 60; j++ {
			for k := 0; k < 60; k++ {
				dp[i][j][k] = -1
			}
		}
	}
	return dfs(m, n, i+1, j+1, N) // 下标取正
}

func dfs(m, n, i, j, k int) int {
	if k < 0 { // 次数够了
		return 0
	}
	if i < 1 || i > m || j < 1 || j > n { // 出界次数+1
		return 1
	}
	if dp[i][j][k] != -1 {
		return dp[i][j][k]
	}
	dp[i][j][k] = 0
	for a := 0; a < 4; a++ { // 上下左右4个方向
		x := i + dx[a]
		y := j + dy[a]
		dp[i][j][k] = (dp[i][j][k] + dfs(m, n, x, y, k-1)) % mod
	}
	return dp[i][j][k]
}
