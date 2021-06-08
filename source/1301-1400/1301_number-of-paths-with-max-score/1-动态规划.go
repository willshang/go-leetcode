package main

func main() {

}

// leetcode1301_最大得分的路径数目
var mod = 1000000007
var dx = []int{1, 0, 1}
var dy = []int{0, 1, 1}

func pathsWithMaxScore(board []string) []int {
	n := len(board)
	dp := make([][][2]int, n) // dp[i][j] 达到坐标i,j的 和最大值 以及 方案数
	for i := 0; i < n; i++ {
		dp[i] = make([][2]int, n)
	}
	dp[n-1][n-1][1] = 1 // 方案数
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if board[i][j] == 'X' {
				continue
			}
			var value int
			if board[i][j] == 'E' || board[i][j] == 'S' {
				value = 0
			} else {
				value = int(board[i][j] - '0')
			}
			for k := 0; k < 3; k++ { // 枚举当前节点下、右、下右节点
				x := i + dx[k]
				y := j + dy[k]
				if x < n && y < n && dp[x][y][1] > 0 {
					sum := dp[x][y][0] + value
					if sum > dp[i][j][0] {
						dp[i][j][0] = sum
						dp[i][j][1] = dp[x][y][1]
					} else if sum == dp[i][j][0] {
						dp[i][j][1] = (dp[i][j][1] + dp[x][y][1]) % mod
					}
				}
			}
		}
	}
	return []int{dp[0][0][0], dp[0][0][1] % mod}
}
