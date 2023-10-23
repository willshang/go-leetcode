package main

func main() {

}

// leetcode1091_二进制矩阵中的最短路径
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
	visited := make(map[[2]int]bool)
	visited[[2]int{0, 0}] = true
	queue := make([][3]int, 0)
	queue = append(queue, [3]int{0, 0, 1})
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		x := node[0]
		y := node[1]
		v := node[2]
		for i := 0; i < 8; i++ {
			newX := x + dx[i]
			newY := y + dy[i]
			if 0 <= newX && newX < n && 0 <= newY && newY < m &&
				grid[newX][newY] == 0 && visited[[2]int{newX, newY}] == false {
				queue = append(queue, [3]int{newX, newY, v + 1})
				visited[[2]int{newX, newY}] = true
				if newX == n-1 && newY == m-1 {
					return v + 1
				}
			}
		}
	}
	return -1
}
