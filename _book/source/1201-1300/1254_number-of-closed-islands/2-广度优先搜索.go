package main

import "fmt"

func main() {
	arr := [][]int{
		{1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 1, 1, 0, 1},
		{1, 0, 1, 0, 1, 0, 1},
		{1, 0, 1, 1, 1, 0, 1},
		{1, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1},
	}
	fmt.Println(closedIsland(arr))
}

var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

func closedIsland(grid [][]int) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				flag := true
				queue := make([][2]int, 0)
				queue = append(queue, [2]int{i, j})
				if i == 0 || j == 0 || i == len(grid)-1 || j == len(grid[i])-1 {
					flag = false
				}
				for len(queue) > 0 {
					node := queue[0]
					queue = queue[1:]
					for k := 0; k < 4; k++ {
						x := dx[k] + node[0]
						y := dy[k] + node[1]
						if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[i]) {
							continue
						}
						if grid[x][y] == 1 {
							continue
						}
						if x == 0 || y == 0 || x == len(grid)-1 || y == len(grid[i])-1 {
							flag = false
						}
						queue = append(queue, [2]int{x, y})
						grid[x][y] = 1
					}
				}
				if flag == true {
					res++
				}
			}
		}
	}
	return res
}
