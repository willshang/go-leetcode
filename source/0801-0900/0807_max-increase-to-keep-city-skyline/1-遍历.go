package main

func main() {

}

// leetcode807_保持城市天际线
func maxIncreaseKeepingSkyline(grid [][]int) int {
	res := 0
	n, m := len(grid), len(grid[0])
	row := make([]int, n)
	col := make([]int, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			row[i] = max(row[i], grid[i][j])
			col[j] = max(col[j], grid[i][j])
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			res = res + min(row[i], col[j]) - grid[i][j]
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
