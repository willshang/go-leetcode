package main

import (
	"sort"
)

func main() {

}

var dx = []int{0, 1, 0, -1}
var dy = []int{1, 0, -1, 0}

func swimInWater(grid [][]int) int {
	n := len(grid)
	return sort.Search(n*n-1, func(target int) bool {
		if target < grid[0][0] {
			return false
		}
		queue := make([][2]int, 0)
		queue = append(queue, [2]int{0, 0})
		visited := make([][]bool, n)
		for i := 0; i < n; i++ {
			visited[i] = make([]bool, n)
		}
		for len(queue) > 0 {
			a, b := queue[0][0], queue[0][1]
			queue = queue[1:]
			if a == n-1 && b == n-1 {
				return true
			}
			for i := 0; i < 4; i++ {
				x, y := a+dx[i], b+dy[i]
				if 0 <= x && x < n && 0 <= y && y < n &&
					visited[x][y] == false && grid[x][y] <= target {
					queue = append(queue, [2]int{x, y})
					visited[x][y] = true
				}
			}
		}
		return false
	})
	return 0
}
