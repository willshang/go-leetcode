package main

func main() {

}

var dx = []int{0, 1, 0, -1}
var dy = []int{1, 0, -1, 0}
var n, m int
var arr [][]int

func longestIncreasingPath(matrix [][]int) int {
	n, m = len(matrix), len(matrix[0])
	arr = make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, m)
	}
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			res = max(res, dfs(matrix, i, j))
		}
	}
	return res
}

func dfs(matrix [][]int, i, j int) int {
	if arr[i][j] != 0 {
		return arr[i][j]
	}
	arr[i][j]++ // 长度为1
	for k := 0; k < 4; k++ {
		x, y := i+dx[k], j+dy[k]
		if 0 <= x && x < n && 0 <= y && y < m && matrix[x][y] > matrix[i][j] {
			arr[i][j] = max(arr[i][j], dfs(matrix, x, y)+1)
		}
	}
	return arr[i][j]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
