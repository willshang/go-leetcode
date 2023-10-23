package main

func main() {

}

// leetcode417_太平洋大西洋水流问题
var dx = []int{0, 1, 0, -1}
var dy = []int{1, 0, -1, 0}
var n, m int

func pacificAtlantic(heights [][]int) [][]int {
	res := make([][]int, 0)
	n, m = len(heights), len(heights[0])
	A, B := make([][]bool, n), make([][]bool, n)
	for i := 0; i < n; i++ {
		A[i], B[i] = make([]bool, m), make([]bool, m)
	}
	for i := 0; i < n; i++ { // 枚举左右两边往中间走
		dfs(heights, A, i, 0) // 最左边（同上边）走到A
		dfs(heights, B, i, m-1)
	}
	for j := 0; j < m; j++ { // 枚举上下两边往中间走
		dfs(heights, A, 0, j) // 最上边（同左边）走到A
		dfs(heights, B, n-1, j)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if A[i][j] == true && B[i][j] == true {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}

func dfs(heights [][]int, visited [][]bool, i, j int) {
	visited[i][j] = true
	for k := 0; k < 4; k++ {
		x, y := dx[k]+i, dy[k]+j
		if 0 <= x && x < n && 0 <= y && y < m &&
			heights[x][y] >= heights[i][j] && visited[x][y] == false {
			dfs(heights, visited, x, y)
		}
	}
}
