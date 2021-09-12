package main

func main() {

}

// 顺时针：上右下左
var dx = []int{0, 1, 0, -1}
var dy = []int{1, 0, -1, 0}
var res [][]int
var a, b, c, d int

func findFarmland(land [][]int) [][]int {
	res = make([][]int, 0)
	n, m := len(land), len(land[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if land[i][j] == 1 {
				a, b, c, d = i, j, i, j
				land[i][j] = 2
				dfs(land, i, j)
				res = append(res, []int{a, b, c, d})
			}
		}
	}
	return res
}

func dfs(land [][]int, i, j int) {
	a, b = min(a, i), min(b, j)
	c, d = max(c, i), max(d, j)
	for k := 0; k < 4; k++ {
		x, y := i+dx[k], j+dy[k]
		if 0 <= x && x < len(land) && 0 <= y && y < len(land[0]) && land[x][y] == 1 {
			land[x][y] = 2
			dfs(land, x, y)
		}
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
