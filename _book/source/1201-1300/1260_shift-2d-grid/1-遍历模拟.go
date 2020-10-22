package main

import "fmt"

func main() {
	arr := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(shiftGrid(arr, 1))
}

func shiftGrid(grid [][]int, k int) [][]int {
	for i := 0; i < k; i++ {
		temp := make([][]int, len(grid))
		for i := 0; i < len(temp); i++ {
			temp[i] = make([]int, len(grid[i]))
		}
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i])-1; j++ {
				temp[i][j+1] = grid[i][j]
			}
		}
		for i := 0; i < len(grid)-1; i++ {
			temp[i+1][0] = grid[i][len(grid[0])-1]
		}
		temp[0][0] = grid[len(grid)-1][len(grid[0])-1]
		grid = temp
	}
	return grid
}
