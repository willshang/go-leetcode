package main

func main() {

}

// leetcode576_出界的路径数
var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}
var mod = 1000000007

func findPaths(m int, n int, N int, i int, j int) int {
	dp := [60][60][60]int{}
	for k := 1; k <= N; k++ {
		for a := 0; a < m; a++ {
			for b := 0; b < n; b++ {
				for i := 0; i < 4; i++ {
					x := a + dx[i]
					y := b + dy[i]
					if x < 0 || x >= m || y < 0 || y >= n { // 出界次数+1
						dp[a][b][k]++
					} else {
						dp[a][b][k] = (dp[a][b][k] + dp[x][y][k-1]) % mod
					}
				}
			}
		}
	}
	return dp[i][j][N] % mod
}
