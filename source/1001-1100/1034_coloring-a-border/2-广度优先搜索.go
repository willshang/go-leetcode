package main

import "fmt"

func main() {
	fmt.Println(colorBorder([][]int{
		{1, 2, 2},
		{2, 3, 2},
	}, 0, 1, 3))
}

var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

func colorBorder(grid [][]int, r0 int, c0 int, color int) [][]int {
	n, m := len(grid), len(grid[0])
	targetColor := grid[r0][c0]
	visited := make([]bool, n*m)
	queue := make([][2]int, 0)
	queue = append(queue, [2]int{r0, c0})
	visited[r0*m+c0] = true
	for len(queue) > 0 {
		a, b := queue[0][0], queue[0][1]
		queue = queue[1:]
		for i := 0; i < 4; i++ {
			x := a + dx[i]
			y := b + dy[i]
			if 0 <= x && x < n && 0 <= y && y < m {
				if visited[x*m+y] == true { // 访问过
					continue
				}
				if grid[x][y] != targetColor { // 颜色不同，边界
					grid[a][b] = color
				} else { // 颜色相同
					visited[x*m+y] = true
					queue = append(queue, [2]int{x, y})
				}
			} else {
				grid[a][b] = color // 边界
			}
		}
	}
	return grid
}
