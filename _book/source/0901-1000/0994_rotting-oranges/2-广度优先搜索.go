package main

import "fmt"

func main() {
	arr := [][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	}
	fmt.Println(orangesRotting(arr))
}

// leetcode994_腐烂的橘子
// 上、右、下、左
var dx = []int{0, 1, 0, -1}
var dy = []int{1, 0, -1, 0}

func orangesRotting(grid [][]int) int {
	queue := make([][]int, 0)
	count := 0
	times := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			} else if grid[i][j] == 1 {
				count = count + 1
			}
		}
	}
	for len(queue) > 0 {
		times++
		length := len(queue)
		for i := 0; i < length; i++ {
			for j := 0; j < 4; j++ {
				x := queue[i][0] + dx[j]
				y := queue[i][1] + dy[j]
				if x >= 0 && x < len(grid) &&
					y >= 0 && y < len(grid[0]) && grid[x][y] == 1 {
					grid[x][y] = 2
					queue = append(queue, []int{x, y})
					count--
				}
			}
		}
		queue = queue[length:]
		if len(queue) == 0 {
			times--
		}
	}
	if count > 0 {
		return -1
	}
	return times
}
