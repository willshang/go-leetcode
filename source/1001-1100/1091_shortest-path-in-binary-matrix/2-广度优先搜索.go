package main

func main() {

}

var dx = []int{-1, -1, -1, 0, 0, 1, 1, 1}
var dy = []int{-1, 0, 1, -1, 1, -1, 0, 1}

func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] == 1 {
		return -1
	}
	n, m := len(grid), len(grid[0])
	if grid[n-1][m-1] == 1 {
		return -1
	}
	if n == 1 && m == 1 {
		return 1
	}
	queue := make([]int, 0)
	queue = append(queue, 0)
	grid[0][0] = 1
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		x := node / m
		y := node % m
		for i := 0; i < 8; i++ {
			newX := x + dx[i]
			newY := y + dy[i]
			if 0 <= newX && newX < n && 0 <= newY && newY < m && grid[newX][newY] == 0 {
				queue = append(queue, newX*m+newY)
				grid[newX][newY] = grid[x][y] + 1
				if newX == n-1 && newY == m-1 {
					return grid[n-1][m-1]
				}
			}
		}
	}
	return -1
}
