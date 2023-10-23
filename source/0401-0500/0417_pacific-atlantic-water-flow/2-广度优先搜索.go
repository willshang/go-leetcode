package main

func main() {

}

var dx = []int{0, 1, 0, -1}
var dy = []int{1, 0, -1, 0}
var queue [][2]int
var n, m int

func pacificAtlantic(heights [][]int) [][]int {
	res := make([][]int, 0)
	n, m = len(heights), len(heights[0])
	queue = make([][2]int, 0)
	A, B := make([][]bool, n), make([][]bool, n)
	for i := 0; i < n; i++ {
		A[i], B[i] = make([]bool, m), make([]bool, m)
	}
	for i := 0; i < n; i++ { // 枚举左右两边往中间走
		queue = append(queue, [2]int{i, 0})
		A[i][0] = true
		bfs(heights, A) // 最左边（同上边）走到A
		queue = append(queue, [2]int{i, m - 1})
		B[i][m-1] = true
		bfs(heights, B)
	}
	for j := 0; j < m; j++ { // 枚举上下两边往中间走
		queue = append(queue, [2]int{0, j})
		A[0][j] = true
		bfs(heights, A) // 最上边（同左边）走到A
		queue = append(queue, [2]int{n - 1, j})
		B[n-1][j] = true
		bfs(heights, B)
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

func bfs(heights [][]int, visited [][]bool) {
	for len(queue) > 0 {
		i, j := queue[0][0], queue[0][1]
		queue = queue[1:]
		for k := 0; k < 4; k++ {
			x, y := dx[k]+i, dy[k]+j
			if 0 <= x && x < n && 0 <= y && y < m &&
				heights[x][y] >= heights[i][j] && visited[x][y] == false {
				visited[x][y] = true
				queue = append(queue, [2]int{x, y})
			}
		}
	}
}
