package main

import "sort"

func main() {

}

var dx = []int{0, 1, 0, -1}
var dy = []int{1, 0, -1, 0}

func minimumEffortPath(heights [][]int) int {
	n, m := len(heights), len(heights[0])
	return sort.Search(1000000, func(target int) bool {
		queue := make([][2]int, 0)
		queue = append(queue, [2]int{0, 0})
		visited := make([][]bool, n)
		for i := 0; i < n; i++ {
			visited[i] = make([]bool, m)
		}
		for len(queue) > 0 {
			a, b := queue[0][0], queue[0][1]
			queue = queue[1:]
			if a == n-1 && b == m-1 {
				return true
			}
			for i := 0; i < 4; i++ {
				x, y := a+dx[i], b+dy[i]
				if 0 <= x && x < n && 0 <= y && y < m &&
					visited[x][y] == false && abs(heights[a][b]-heights[x][y]) <= target {
					queue = append(queue, [2]int{x, y})
					visited[x][y] = true
				}
			}
		}
		return false
	})
	return 0
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
