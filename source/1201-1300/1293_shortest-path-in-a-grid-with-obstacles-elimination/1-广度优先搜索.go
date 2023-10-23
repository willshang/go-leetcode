package main

func main() {

}

// leetcode1293_网格中的最短路径
var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

func shortestPath(grid [][]int, k int) int {
	n, m := len(grid), len(grid[0])
	if n == 1 && m == 1 {
		return 0
	}
	k = min(k, n+m-3) // 缩小k的范围
	visited := make([][][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([][]bool, m)
		for j := 0; j < m; j++ {
			visited[i][j] = make([]bool, k+1)
		}
	}
	count := 1
	queue := make([][3]int, 0)
	queue = append(queue, [3]int{0, 0, k})
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			a, b, c := queue[i][0], queue[i][1], queue[i][2]
			for j := 0; j < 4; j++ {
				x, y := a+dx[j], b+dy[j]
				if 0 <= x && x < n && 0 <= y && y < m {
					if grid[x][y] == 0 && visited[x][y][c] == false {
						if x == n-1 && y == m-1 { // 走到终点
							return count
						}
						queue = append(queue, [3]int{x, y, c})
						visited[x][y][c] = true
					} else if grid[x][y] == 1 && c > 0 && visited[x][y][c-1] == false {
						queue = append(queue, [3]int{x, y, c - 1})
						visited[x][y][c-1] = true
					}
				}
			}
		}
		queue = queue[length:]
		count++
	}
	return -1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
