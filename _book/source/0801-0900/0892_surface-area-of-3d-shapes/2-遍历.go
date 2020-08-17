package main

import "fmt"

func main() {
	arr := [][]int{{2, 2, 2}, {2, 1, 2}, {2, 2, 2}}
	fmt.Println(surfaceArea(arr))
}

// 上、右、下、左
var dx = []int{0, 1, 0, -1}
var dy = []int{1, 0, -1, 0}

func surfaceArea(grid [][]int) int {
	sum := 0
	for i, rows := range grid {
		for j, _ := range rows {
			sum = sum + grid[i][j]*6
			if grid[i][j] > 1 {
				sum = sum - (grid[i][j]-1)*2
			}
			for k := 0; k < 4; k++ {
				x, y := i+dx[k], j+dy[k]
				if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) {
					sum = sum - min(grid[i][j], grid[x][y])
				}
			}
		}
	}
	return sum
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
