package main

import "fmt"

func main() {
	arr := [][]int{
		{1, 0, 1},
		{0, 0, 0},
		{1, 0, 1},
	}
	fmt.Println(maxDistance(arr))
}

// leetcode1162_地图分析
var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

func maxDistance(grid [][]int) int {
	queue := make([][2]int, 0)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}
	if len(queue) == 0 || len(queue) == len(grid)*len(grid[0]) {
		return -1
	}
	res := -1
	for len(queue) > 0 {
		res++
		length := len(queue)
		for i := 0; i < length; i++ {
			x1 := queue[i][0]
			y1 := queue[i][1]
			for i := 0; i < 4; i++ {
				x := x1 + dx[i]
				y := y1 + dy[i]
				if 0 <= x && x < len(grid) && 0 <= y && y < len(grid[0]) && grid[x][y] == 0 {
					queue = append(queue, [2]int{x, y})
					grid[x][y] = 2
				}
			}
		}
		queue = queue[length:]
	}
	return res
}
